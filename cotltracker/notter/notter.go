package notter

import "player/cotltracker/inputter"

// Notter - Позволяет воспроизводить простые звуки
type Notter struct {
	input *inputter.Inputer
}

// New - Создаёт новый экземпляр нотного блока
func New(input *inputter.Inputer) *Notter {
	return &Notter{
		input: input,
	}
}
