package chorder

import (
	"fmt"
)

func (chorder *Chorder) PlayChord(chord string) error {
	chords := map[string]*Chord{
		"Am": NewChord("C4 E4 A4", -4),
		"Bm": NewChord("D4 B4 F5", -4),
		"Em": NewChord("E4 G4 B4", -4),
		"G":  NewChord("D4 G4 B4", -4),
		"A":  NewChord("D4 E4 A4", -4),
	}

	c, exists := chords[chord]
	if !exists {
		return fmt.Errorf("unknown chord: %s", chord)
	}

	// Воспроизведение всех нот
	for _, note := range c.notes {
		if err := chorder.notter.PlayNote(note, c.octaveShift); err != nil {
			return err
		}
	}

	return nil
}
