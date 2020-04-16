package tapper

import (
	"github.com/kunaldawn/goandroid/input"
)

// Tapper - Управляет вводом на смартфоне
type Tapper struct {
	screen input.TouchScreen
}

// New - Создаёт новый экземпляр управления вводом
func New(screen input.TouchScreen) *Tapper {
	return &Tapper{
		screen: screen,
	}
}
