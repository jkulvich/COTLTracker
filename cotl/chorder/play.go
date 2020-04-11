package chorder

import (
	"fmt"
)

//https://gitarre.ru/akkord-gm

// PlayChord - Воспроизводит аккорд
func (chorder *Chorder) PlayChord(chord string) error {
	chords := map[string]*Chord{
		"A": NewChord("A4 D4 E4", -4),
		"Am": NewChord("A4 C4 E4", -4),

		"B": NewChord("B4 E4 F5", -4),
		"Bm": NewChord("B4 D4 F5", -4),

		"C": NewChord("C4 E4 G4", -4),
		"Cm": NewChord("C5 E4 G5", -4),

		"D": NewChord("D4 F5 A4", -4),
		"Dm": NewChord("D4 F4 A4", -4),

		"E": NewChord("E4 G5 B4", -4),
		"Em": NewChord("E5 G5 B4", -4),

		"F": NewChord("F4 A5 C4", -4),
		"Fm": NewChord("F5 A5 C4", -4),

		"G":  NewChord("G4 B4 D4", -4),
		"Gm":  NewChord("G5 B4 D5", -4),
	}

	c, exists := chords[chord]
	if !exists {
		return fmt.Errorf("unknown chord: %s", chord)
	}

	// Воспроизведение всех нот
	for _, note := range c.notes {
		if err := chorder.notter.PlayNote(note, c.octaveShift * 7); err != nil {
			return err
		}
	}

	return nil
}
