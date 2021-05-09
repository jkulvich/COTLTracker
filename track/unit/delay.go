package unit

import (
	"fmt"
	"regexp"
	"strconv"
)

// Delay - Delay in ticks
type Delay uint8

// NewDelay - Make delay from any known representation string.
// Numbers like: 200 500 1000, in this case it will be convert by formula ticks = round(ms/100).
// tN like: t1 t2 t3, number after "t" will be used as a ticks' count.
// Dashes: - --- ----, count of dashes will be used as a ticks' count.
func NewDelay(str string) (Delay, error) {
	// Number notation
	if d, err := strconv.Atoi(str); err == nil {
		return Delay(d / 100), nil
	}
	// Dash notation
	if regexp.MustCompile(`-+`).MatchString(str) {
		return Delay(len(str)), nil
	}
	// tN notation
	if matches := regexp.MustCompile(`t(\d+)`).FindStringSubmatch(str); len(matches) == 2 {
		t, _ := strconv.Atoi(matches[1])
		return Delay(t), nil
	}
	return 0, fmt.Errorf("unknown delay type notation: %s", str)
}
