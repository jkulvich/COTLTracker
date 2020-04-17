package cotl

import (
	"fmt"
	"math"
	"player/cotl/block"
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
	track := Track{
		blocks: make(block.Blocks, 0),
		shift:  0,
		timing: 200,
	}

	// Применяем настройки из музыкального файла
	if err := track.applyConfigComments(stave); err != nil {
		return nil, err
	}

	stave = track.removeComments(stave)
	stave = strings.ReplaceAll(stave, "\n", " ")
	runes := strings.Split(stave, " ")

	for _, r := range runes {
		r = strings.TrimSpace(r)
		if len(r) > 0 {
			blocks, err := block.NewBlocks(r, track.timing)
			if err != nil {
				return nil, err
			}
			for _, blk := range blocks {
				track.blocks = append(track.blocks, blk)
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
