package cotl

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
)

type trackBlocks []*trackBlock

// Track - Содержит настройки композиции и последовательность музыкальных блоков
type Track struct {
	blocks      trackBlocks
	shift       int //< Пользовательский сдвиг
	octaveShift int //< Автоматически определённый сдвиг для октавы
	timing      int //< Множитель для задержек тире и t
}

// NewTrack - Парсит нотное представление и представляет его в виде трека
func NewTrack(stave string) (*Track, error) {
	track := Track{
		blocks:      make(trackBlocks, 0),
		shift:       0,
		octaveShift: 0,
		timing:      200,
	}

	// Применяем настройки из музыкального файла
	track.applyConfigComments(stave)

	stave = track.removeComments(stave)
	stave = strings.ReplaceAll(stave, "\n", " ")
	runes := strings.Split(stave, " ")

	for _, r := range runes {
		r = strings.TrimSpace(r)
		if len(r) > 0 {
			block, err := track.newTrackBlock(r)
			if err != nil {
				return nil, err
			}
			track.blocks = append(track.blocks, block)
		}
	}

	// Пытаемся автоматически определить сдвиг ноты по октаве
	octaveShift, err := track.detectOctaveShift(track.blocks)
	if err != nil {
		return nil, err
	}
	track.octaveShift = octaveShift

	return &track, nil
}

// SetTiming - Устанавливает тайминг для композиции
func (track *Track) SetTiming(timing int) *Track {
	track.timing = timing
	return track
}

// SetShift - Устанавливает сдвиг для композиции
func (track *Track) SetShift(shift int) *Track {
	track.shift = shift
	return track
}

// removeComments - Удаляет все строки начинающиеся с #
func (*Track) removeComments(stave string) string {
	lines := strings.Split(stave, "\n")
	nlines := make([]string, 0, len(lines))

	for _, line := range lines {
		if !strings.HasPrefix(line, "#") {
			nlines = append(nlines, line)
		}
	}

	return strings.Join(nlines, "\n")
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
func (*Track) detectOctaveShift(blocks trackBlocks) (int, error) {
	minOct := math.MaxInt32

	for _, block := range blocks {
		if block.typ == trackBlockNote {

			res := regexp.MustCompile(`([ABCDEFG])(\d*)`).FindStringSubmatch(strings.ToUpper(block.note))
			if len(res) != 3 {
				return 0, fmt.Errorf("unexpected note form: %s", block.note)
			}

			octave, err := strconv.Atoi(res[2])
			if err != nil {
				return 0, fmt.Errorf("unexpected note form: %s", block.note)
			}

			if octave < minOct {
				minOct = octave
			}

		}
	}

	// Если значение не сменилось - в композиции нет нот, не двигаем
	if minOct == math.MaxInt32 {
		minOct = 0
	}

	return -minOct * 7, nil
}
