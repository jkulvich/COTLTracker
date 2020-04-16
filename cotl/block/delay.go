package block

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
)

type Delay struct {
	timing int
	delay  int
}

// NewDelay - Создаёт новый блок задержки получая один из следующих форматов:
// 200 500 1000 - Простое число воспринимается как миллисекунды
// t1 t2 t4 - Воспринимается как timing*N
// - --- ----- - Воспринимается как len(str)*N
func NewDelay(delay string, timing int) (Delay, error) {
	calculatedDelay := 0

	if d, err := strconv.Atoi(delay); err == nil {
		calculatedDelay = d
	} else if regexp.MustCompile(`-+`).MatchString(delay) {
		calculatedDelay = int(len(delay)) * timing
	} else if matches := regexp.MustCompile(`t(\d+)`).FindStringSubmatch(delay); len(matches) == 2 {
		t, _ := strconv.Atoi(matches[1])
		calculatedDelay = int(t) * timing
	} else {
		return Delay{}, fmt.Errorf("unknown delay type %s", delay)
	}

	return Delay{
		timing: timing,
		delay:  calculatedDelay,
	}, nil
}

// GetDelay - Возвращает задержку которая должна произойти после выполнения блока
func (delay Delay) GetDelay() int {
	return delay.delay
}

// String - Представление задержки в виде тире или числом если не выходит остатка без деления
func (delay Delay) String() string {
	count := float64(delay.delay) / float64(delay.timing)
	if count == math.Trunc(count) {
		return strings.Repeat("-", int(count))
	}
	return strconv.Itoa(int(delay.delay))
}

// GetBlocks - Возвращает музыкальный блок
func (delay Delay) GetBlocks() []Block {
	return []Block{{Delay: delay.delay}}
}
