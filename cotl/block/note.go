package block

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// Note - Объект ноты включающий тон и октаву
type Note struct {
	Octave int
	Tone   uint8
}

// NewNote - Создаёт новый объект ноты
func NewNote(note string) (Note, error) {
	reg := regexp.MustCompile(`(?i)$?([A-H])(\d*)`)
	if !reg.MatchString(note) {
		return Note{}, fmt.Errorf("can't parse %s as block, must be like [A-H]\\d", note)
	}

	// Парсим тон и октаву
	res := reg.FindStringSubmatch(strings.ToUpper(note))
	tone := res[1][0]
	octave, err := strconv.Atoi(res[2])
	if err != nil {
		return Note{}, fmt.Errorf("can't parse block %s number", note)
	}

	// Защита от упоротой ноты H
	if tone == 'H' {
		tone = 'B'
	}

	return Note{
		Octave: octave,
		Tone:   tone,
	}, nil
}

// String - Получение текстового представления ноты
func (note Note) String() string {
	return fmt.Sprintf("%c%d", note.Tone, note.Octave)
}

// Hash - Числовое представление ноты
func (note Note) Hash() int {
	// Карта замены названия тонов для корректной калькуляции
	rpl := map[uint8]uint8{
		'C': 'A',
		'D': 'B',
		'E': 'C',
		'F': 'D',
		'G': 'E',
		'A': 'F',
		'B': 'G',
	}
	return int(rpl[note.Tone]-'A') + note.Octave*7
}

// HashRestore - Восстанавливает ноту из её хэша
func (note *Note) HashRestore(hash int) error {
	if hash < 0 {
		return fmt.Errorf("negative hash %d", hash)
	}

	rplBack := map[uint8]uint8{
		'A': 'C',
		'B': 'D',
		'C': 'E',
		'D': 'F',
		'E': 'G',
		'F': 'A',
		'G': 'B',
	}

	octave := hash / 7
	tone, exists := rplBack[uint8(int('A')+hash%7)]
	if !exists {
		return fmt.Errorf("incorrect Tone %d", hash)
	}

	if octave < 0 {
		return fmt.Errorf("negative Octave %d", octave)
	}

	note.Octave = octave
	note.Tone = tone

	return nil
}

// ShiftNote - Сдвигает ноту по звучанию на следующую ноту (+7 - сдвиг на октаву)
func (note *Note) ShiftNote(shift int) error {
	hash := note.Hash()
	hash += shift
	if err := note.HashRestore(hash); err != nil {
		return fmt.Errorf("can't shift: %s", err)
	}
	return nil
}

// GetBlocks - Возвращает музыкальный блок
func (note Note) GetBlocks() []Block {
	return []Block{{Note: &note}}
}
