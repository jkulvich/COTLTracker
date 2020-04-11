package notter

import (
	"fmt"
)

type tone map[uint8][]int
type octave map[int]tone

// playSound - Воспроизводит отдельный звук по тону (A-G) и октаве (0-2)
func (nott *Notter) playSound(ton uint8, oct int) error {
	octaves := octave{
		0: tone{
			'C': []int{0, 0},
			'D': []int{1, 0},
			'E': []int{2, 0},
			'F': []int{3, 0},
			'G': []int{4, 0},
			'A': []int{0, 1},
			'B': []int{1, 1},
		},
		1: tone{
			'C': []int{2, 1},
			'D': []int{3, 1},
			'E': []int{4, 1},
			'F': []int{0, 2},
			'G': []int{1, 2},
			'A': []int{2, 2},
			'B': []int{3, 2},
		},
		2: tone{
			'C': []int{4, 2},
		},
	}

	currentOctave, exists := octaves[oct]
	if !exists {
		return fmt.Errorf("incorrect octave, expected 0-2, got %d", oct)
	}

	currentTone, exists := currentOctave[ton]
	if !exists {
		return fmt.Errorf("incorrect tone in current octave, got tone %c%d", ton, oct)
	}

	// Позиция клавиши для нажатия
	col := currentTone[0]
	row := currentTone[1]

	if err := nott.input.TapButton(col, row); err != nil {
		return err
	}

	return nil
}
