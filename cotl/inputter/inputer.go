package inputter

import (
	"github.com/kunaldawn/goandroid/input"
)

// Inputer - Управляет вводом на смартфоне
type Inputer struct {
	screen input.TouchScreen
}

// New - Создаёт новый экземпляр управления вводом
func New(screen input.TouchScreen) *Inputer {
	return &Inputer{
		screen: screen,
	}
}
