package cotl

import (
	"fmt"
	"regexp"
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
func (track *Track) newTrackNote(note string) *trackBlock {
	return &trackBlock{
		typ:  trackBlockNote,
		note: note,
	}
}

// newTrackChord - Новый блок характеризующий аккорд
func (track *Track) newTrackChord(chord string) *trackBlock {
	return &trackBlock{
		typ:   trackBlockChord,
		chord: chord,
	}
}

// newTrackDelay - Новый блок характеризующий задержку в мс
func (track *Track) newTrackDelay(delay int) *trackBlock {
	return &trackBlock{
		typ:   trackBlockDelay,
		delay: delay,
	}
}

// newTrackBlock - Определяет тип блока по краткой текстовой записи
func (track *Track) newTrackBlock(rune string) (*trackBlock, error) {
	// Ноты могут вылгядеть тае (C4 D5 A2 опционально знак $ перед нотой)
	// Задержка в мс - просто числа (200 500 1000)
	// Всё остальное - аккорды (Am Dm E Bm)

	// Попытка определить ноту
	if regexp.MustCompile(`(?i)\$?[A-H][1-6]`).MatchString(rune) {
		rune = strings.TrimPrefix(rune, "$")
		return track.newTrackNote(rune), nil
	}

	// Попытка определить задержку для:
	// Указания через время
	// Указания через тире
	// Указания через тайминг
	if delay, err := strconv.Atoi(rune); err == nil {

		return track.newTrackDelay(delay), nil

	} else if regexp.MustCompile(`-+`).MatchString(rune) {

		delay := int(len(rune)) * track.timing
		return track.newTrackDelay(delay), nil

	} else if matches := regexp.MustCompile(`t(\d+)`).FindStringSubmatch(rune); len(matches) == 2 {

		t, _ := strconv.Atoi(matches[1])
		delay := int(t) * track.timing
		return track.newTrackDelay(delay), nil

	}

	// Попытка определить аккорд
	if regexp.MustCompile(`(?i)[A-H].*`).MatchString(rune) {
		return track.newTrackChord(rune), nil
	}

	return nil, fmt.Errorf("can't recognize the rune: %s", rune)
}
