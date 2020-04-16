package block

import "fmt"

// chord - Хранит аккорд
type Chord struct {
	chord string
	notes []Note
}

// NewChord - Создаёт новый аккорд если такой известен системе
func NewChord(chord string) (Chord, error) {
	var chords = map[string][]string{
		"A":  []string{"A0", "D0", "E0"},
		"Am": []string{"A0", "C0", "E0"},
		"B":  []string{"B0", "E0", "F1"},
		"Bm": []string{"B0", "D0", "F1"},
		"C":  []string{"C0", "E0", "G0"},
		"Cm": []string{"C1", "E0", "G1"},
		"D":  []string{"D0", "F1", "A0"},
		"Dm": []string{"D0", "F0", "A0"},
		"E":  []string{"E0", "G1", "B0"},
		"Em": []string{"E1", "G1", "B0"},
		"F":  []string{"F0", "A1", "C0"},
		"Fm": []string{"F1", "A1", "C0"},
		"G":  []string{"G0", "B0", "D0"},
		"Gm": []string{"G1", "B0", "D1"},
	}

	notesNames, exists := chords[chord]
	if !exists {
		return Chord{}, fmt.Errorf("unknown chord %s", chord)
	}

	notes := make([]Note, 0, len(notesNames))
	for _, n := range notesNames {
		note, err := NewNote(n)
		if err != nil {
			return Chord{}, fmt.Errorf("can't get chord notes %s", chord)
		}
		notes = append(notes, note)
	}

	return Chord{
		chord: chord,
		notes: notes,
	}, nil
}

// String - Текстовое представление аккорда
func (chord Chord) String() string {
	return chord.chord
}

// GetNotes - Возвращает воспроизводимые ноты, нужно для интерфейса
func (chord Chord) GetNotes() []Note {
	return chord.notes
}

// GetBlocks - Возвращает музыкальные блоки
func (chord Chord) GetBlocks() []Block {
	blocks := make([]Block, 0, len(chord.notes))
	for _, note := range chord.notes {
		for _, blk := range note.GetBlocks() {
			blocks = append(blocks, blk)
		}
	}
	return blocks
}
