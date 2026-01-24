package cotl

import (
	"encoding/json"
	"fmt"
	"math"
	"player/cotl/block"
	"sort"
	"strconv"
	"strings"
)

// Track - Содержит настройки композиции и последовательность музыкальных блоков
type Track struct {
	blocks block.Blocks
	shift  int //< Пользовательский сдвиг
	timing int //< Множитель для задержек тире и t
}

// NewTrack - Парсит нотное представление и представляет его в виде трека
func NewTrack(stave string) (*Track, error) {
	if json.Valid([]byte(stave)) {
		if track, err := NewTrackFromSkyMusicSheet(stave); err == nil {
			return track, nil
		}
	}

	return NewTrackFromCOTLTrackerFormat(stave)
}

// NewTrackFromSkyMusicSheet - Sky Music Sheet Maker format.
// This is a very-very crunchy parser, but it is quick hack to work with initially unsupported format :)
func NewTrackFromSkyMusicSheet(stave string) (*Track, error) {
	var parties []struct {
		CotlTrackerSpeedup float32 `json:"cotltrackerSpeedup"`
		Name               string  `json:"name"`
		Author             string  `json:"author"`
		TanscribedBy       string  `json:"transcribedBy"`
		IsComposed         bool    `json:"isComposed"`
		BpM                int     `json:"bpm"`
		BitsPerPage        int     `json:"bitsPerPage"`
		PitchLevel         int     `json:"pitchLevel"`
		IsEncrypted        bool    `json:"isEncrypted"`
		SongNotes          []struct {
			Time int    `json:"time"`
			Key  string `json:"key"`
		} `json:"songNotes"`
	}

	if err := json.Unmarshal([]byte(stave), &parties); err != nil {
		return nil, err
	}

	part := parties[0]

	// sort notes to be consistent
	sort.Slice(part.SongNotes, func(i, j int) bool {
		return part.SongNotes[i].Time < part.SongNotes[j].Time
	})

	track := Track{
		blocks: make(block.Blocks, 0),
		shift:  0,
		timing: 200,
	}

	if part.CotlTrackerSpeedup == 0 {
		// delay compensation from ADB input
		part.CotlTrackerSpeedup = 1.11
	}

	// iterate over notes and create blocks
	previousTime := 0
	posCounter := 0
	for iNote, note := range part.SongNotes {
		if previousTime < note.Time {
			delayTime := note.Time - previousTime
			delayTime = int(float32(delayTime) / part.CotlTrackerSpeedup)
			delay, err := block.NewDelay(strconv.Itoa(delayTime), 0)
			if err != nil {
				return nil, err
			}
			// skip initial delay
			if iNote > 0 {
				posCounter++
				delayBlocks := delay.GetBlocks()
				delayBlocks[0].LineNum = posCounter
				track.blocks = append(track.blocks, delayBlocks...)
			}
			previousTime = note.Time
		}
		mapping := map[string]string{
			"1Key0": "C0",
			"1Key1": "D0",
			"1Key2": "E0",
			"1Key3": "F0",
			"1Key4": "G0",

			"1Key5": "A0",
			"1Key6": "B0",
			"1Key7": "C1",
			"1Key8": "D1",
			"1Key9": "E1",

			"1Key10": "F1",
			"1Key11": "G1",
			"1Key12": "A1",
			"1Key13": "B1",
			"1Key14": "C2",

			"2Key0": "C0",
			"2Key1": "D0",
			"2Key2": "E0",
			"2Key3": "F0",
			"2Key4": "G0",

			"2Key5": "A0",
			"2Key6": "B0",
			"2Key7": "C1",
			"2Key8": "D1",
			"2Key9": "E1",

			"2Key10": "F1",
			"2Key11": "G1",
			"2Key12": "A1",
			"2Key13": "B1",
			"2Key14": "C2",
		}
		note, err := block.NewNote(mapping[note.Key])
		if err != nil {
			return nil, err
		}
		posCounter++
		noteBlocks := note.GetBlocks()
		noteBlocks[0].LineNum = posCounter
		track.blocks = append(track.blocks, noteBlocks...)
	}

	return &track, nil
}

// NewTrackFromCOTLTrackerFormat - COTLTracker format
func NewTrackFromCOTLTrackerFormat(stave string) (*Track, error) {
	track := Track{
		blocks: make(block.Blocks, 0),
		shift:  0,
		timing: 200,
	}

	// Применяем настройки из музыкального файла
	if err := track.applyConfigComments(stave); err != nil {
		return nil, err
	}

	for iline, line := range strings.Split(stave, "\n") {
		if strings.HasPrefix(line, "#") {
			continue
		}

		runes := strings.Split(line, " ")

		for irune, r := range runes {
			r = strings.TrimSpace(r)
			if len(r) > 0 {
				blocks, err := block.NewBlocks(r, track.timing, iline, irune)
				if err != nil {
					return nil, err
				}
				for _, blk := range blocks {
					blk.LineNum = iline + 1
					blk.PosNum = irune + 1
					track.blocks = append(track.blocks, blk)
				}
			}
		}
	}

	// Пытаемся автоматически определить сдвиг ноты по октаве
	octaveShift, err := track.detectOctaveShift(track.blocks)
	if err != nil {
		return nil, err
	}

	// Проходим по всем нотам и сдвигаем их
	for i := range track.blocks {
		note := track.blocks[i].Note
		if note != nil {
			if err := note.ShiftNote(octaveShift + track.shift); err != nil {
				return nil, err
			}
		}
	}

	return &track, nil
}

// applyConfigComments - Применяет конфигурацию описанную в управляющий комментариях #!
func (track *Track) applyConfigComments(stave string) error {
	lines := strings.Split(stave, "\n")

	for _, line := range lines {
		if strings.HasPrefix(line, "#!") {
			confLine := strings.TrimPrefix(line, "#!")
			config := strings.Split(confLine, ":")

			switch strings.ToUpper(config[0]) {
			case "TIMING":
				if len(config) != 2 {
					return fmt.Errorf("incorrect config line TIMING, argument expected")
				}
				timing, err := strconv.Atoi(config[1])
				if err != nil {
					return fmt.Errorf("incorrect TIMING: %v", err)
				}
				track.timing = timing
			case "SHIFT":
				if len(config) != 2 {
					return fmt.Errorf("incorrect config line SHIFT, argument expected")
				}
				shift, err := strconv.Atoi(config[1])
				if err != nil {
					return fmt.Errorf("incorrect SHIFT: %v", err)
				}
				track.shift = shift
			}
		}
	}

	return nil
}

// detectOctaveShift - Автоматически определяет сдвиг октавы для нот
func (*Track) detectOctaveShift(blocks block.Blocks) (int, error) {
	minOct := math.MaxInt32

	for _, blk := range blocks {
		if blk.Note != nil && blk.Note.Octave < minOct {
			minOct = blk.Note.Octave
		}
	}

	// Если значение не сменилось - в композиции нет нот, не двигаем
	if minOct == math.MaxInt32 {
		minOct = 0
	}

	return -minOct * 7, nil
}
