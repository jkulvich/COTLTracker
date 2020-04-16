package tapper

import (
	"fmt"
	"player/cotl/block"
	"time"
)

const (
	tapDelay = 20
)

type tone map[uint8][]int
type octave map[int]tone

// tapScreen - Кликает на заданную кооринату экрана
func (input *Tapper) TapPosition(x, y int) error {
	//return input.screen.Tap(x, y)
	go func() {
		input.screen.Swipe(x, y, x, y, 0)
	}()
	<-time.After(time.Millisecond * tapDelay)
	return nil
}

// tapButton - Кликает на указанную кнопку по её позиции в матрице кнопок
func (input *Tapper) TapButton(col, row int) error {
	const zeroX = 730 // Левый верхний угол
	const zeroY = 110

	const shiftX = 190 // Смещение для каждой оси для соседней кнопки
	const shiftY = 190

	const cols = 5 // Количество кнопок
	const rows = 3

	if col < 0 || row < 0 || col >= cols || row >= rows {
		return fmt.Errorf(
			"out of range, got col: %d, row: %d, but zone is col: 0 - %d, row: 0 - %d", col, row, cols-1, rows-1)
	}

	x := zeroX + shiftX*col
	y := zeroY + shiftY*row

	return input.TapPosition(x, y)
}

// TapNote - Кликает на игровую кнопку соответствующую ноте
func (input *Tapper) TapNote(note block.Note) error {
	var octaves = octave{
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
	btn := octaves[note.Octave][note.Tone]
	if len(btn) != 2 {
		return fmt.Errorf("incorrect note tone or octave for play %c%d, use A0-D1", note.Tone, note.Octave)
	}
	return input.TapButton(btn[0], btn[1])
}
