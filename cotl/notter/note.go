package notter

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// PlayNote - Воспроизводит ноту в формате A0, G5, C3 итп, принимает значение транспонирования ноты (-7 сдвиг на октаву)
func (nott *Notter) PlayNote(note string, shift int) error {
	// Транспонируем ноту
	note, err := nott.shiftNote(note, shift)
	if err != nil {
		return err
	}

	// Парсим ноту
	tone, octave, err := nott.parseNote(note)
	return nott.playSound(tone, octave)
}

// shift - Принимает ноту и производит сдвиг по тонам и октавам
func (nott *Notter) shiftNote(note string, shift int) (string, error) {
	tone, octave, err := nott.parseNote(note)
	if err != nil {
		return "", err
	}

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
	rplBack := map[uint8]uint8{
		'A': 'C',
		'B': 'D',
		'C': 'E',
		'D': 'F',
		'E': 'G',
		'F': 'A',
		'G': 'B',
	}

	// Преобразуем в числовое представление
	code := int(rpl[tone]-'A') + octave*7

	// Транспонируем
	code += shift

	// Преобразуем обратно в буквенное представление
	octave = code / 7
	tone = rplBack[uint8(int('A')+code%7)]
	note = fmt.Sprintf("%c%d", tone, octave)

	// Проверка
	if octave < 0 {
		return "", fmt.Errorf("incorrect note octave: %s", note)
	}
	if tone < 'A' || tone > 'G' {
		return "", fmt.Errorf("incorrect note tone: %s", note)
	}

	return note, nil
}

// parseNote - Парсит ноту на тон и октаву
func (nott *Notter) parseNote(note string) (uint8, int, error) {
	reg := regexp.MustCompile(`([A-H])(\d*)`)
	if !reg.MatchString(note) {
		return 0, 0, fmt.Errorf("can't recognize the note: %s", note)
	}

	// Парсим тон и октаву
	res := reg.FindStringSubmatch(strings.ToUpper(note))
	tone := res[1][0]
	octave, _ := strconv.Atoi(res[2])

	// Защита от ноты H
	if tone == 'H' {
		tone = 'B'
	}

	return tone, octave, nil
}
