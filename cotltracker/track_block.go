package cotltracker

import (
	"strconv"
	"strings"
)

type trackBlockType string

const (
	trackBlockNote  trackBlockType = "note"
	trackBlockChord trackBlockType = "chord"
	trackBlockDelay trackBlockType = "delay"
)

// trackBlock - Блок информации о воспроизводимом звуке, аккорде или задержке
type trackBlock struct {
	typ   trackBlockType
	note  string
	chord string
	delay int
}

// newTrackNote - Новый блок характеризующий ноту
func newTrackNote(note string) *trackBlock {
	return &trackBlock{
		typ:  trackBlockNote,
		note: note,
	}
}

// newTrackChord - Новый блок характеризующий аккорд
func newTrackChord(chord string) *trackBlock {
	return &trackBlock{
		typ:   trackBlockChord,
		chord: chord,
	}
}

// newTrackDelay - Новый блок характеризующий задержку в мс
func newTrackDelay(delay int) *trackBlock {
	return &trackBlock{
		typ:   trackBlockDelay,
		delay: delay,
	}
}

// parseTrackBlock - Определяет тип блока по краткой текстовой записи
func parseBlock(rune string) *trackBlock {
	// Все ноты начинаются с $ ($C4 $D5 $A2)
	// Задержка в мс - просто числа (200 500 1000)
	// Всё остальное - аккорды (Am Dm E Bm)

	if strings.HasPrefix(rune, "$") {
		note := strings.TrimPrefix(rune, "$")
		return newTrackNote(note)
	}
	if delay, err := strconv.Atoi(rune); err == nil {
		return newTrackDelay(delay)
	}

	return newTrackChord(rune)
}
