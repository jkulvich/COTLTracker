package notter

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// PlayNote - Воспроизводит ноту в формате A0, G5, C3 итп, принимает корректирующий коэффициент по первой октаве
func (nott *Notter) PlayNote(note string, shift ...int) error {
	res := regexp.MustCompile(`([ABCDEFG])(\d*)`).FindStringSubmatch(strings.ToUpper(note))

	tone := "A"
	octave := 0

	// Разбор ноты на тон и октаву
	if len(res) != 3 {
		return fmt.Errorf("incorrect note, got: %s", note)
	}

	tone = res[1]
	oct, err := strconv.Atoi(res[2])
	if err != nil {
		return err
	}
	octave = oct

	// Применение сдвга октавы
	if len(shift) > 0 {
		octave += shift[0]
	}

	return nott.playSound(tone, octave)
}
