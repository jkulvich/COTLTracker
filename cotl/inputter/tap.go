package inputter

import (
	"fmt"
	"time"
)

const (
	tapDelay = 20
)

// tapScreen - Кликает на заданную кооринату экрана
func (input *Inputer) TapScreen(x, y int) error {
	//return input.screen.Tap(x, y)
	go func() {
		input.screen.Swipe(x, y, x, y, 0)
	}()
	<-time.After(time.Millisecond * tapDelay)
	return nil
}

// tapButton - Кликает на указанную кнопку по её позиции в матрице кнопок
func (input *Inputer) TapButton(col, row int) error {
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

	return input.TapScreen(x, y)
}
