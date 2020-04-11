package chorder

import "strings"

// Chord - Описание структуры ноты
type Chord struct {
	notes       []string
	octaveShift int
}

// NewChord - Новое описание аккорда
func NewChord(notes string, octaveShift int) *Chord {
	return &Chord{
		notes:       strings.Split(notes, " "),
		octaveShift: octaveShift,
	}
}
