package android

import (
	"fmt"
	"time"
)

// Swipe - It makes a swipe action from x1/y1 to x2/y2 for 40ms by default
func (dev *Android) Swipe(x1, y1, x2, y2 int, delay ...int) error {
	const safeMinDelay = 40

	d := safeMinDelay
	if len(delay) > 0 {
		d = delay[0]
	}

	cmd := dev.SwipeCmd(x1, y1, x2, y2, d)
	if _, err := dev.Exec(cmd); err != nil {
		return err
	}
	return nil
}

// SwipeCmd - Returns swipe command for shell
func (dev *Android) SwipeCmd(x1, y1, x2, y2 int, delay ...int) string {
	const safeMinDelay = 40

	d := safeMinDelay
	if len(delay) > 0 {
		d = delay[0]
	}

	return fmt.Sprintf("input swipe %d %d %d %d %d", x1, y1, x2, y2, d)
}

// Tap - It makes a tap in specific coordinates
func (dev *Android) Tap(x, y int) error {
	cmd := fmt.Sprintf("input tap %d %d", x, y)
	if _, err := dev.Exec(cmd); err != nil {
		return err
	}
	return nil
}

// Key - It makes a key event with specific key code
func (dev *Android) Key(code int) error {
	cmd := fmt.Sprintf("input keyevent %d", code)
	if _, err := dev.Exec(cmd); err != nil {
		return err
	}
	return nil
}

// FastTap - A little hack for ADB Tap to speed improve
func (dev *Android) FastTap(x, y int, delay ...int) {
	d := dev.MinTapDelay
	if len(delay) > 0 {
		d = delay[0]
	}

	go func() {
		_ = dev.Swipe(x, y, x, y)
	}()
	<-time.After(time.Millisecond * time.Duration(d))
}

// FastTapCmd - Returns fast tap shell command
func (dev *Android) FastTapCmd(x, y int, delay ...int) string {
	d := dev.MinTapDelay
	if len(delay) > 0 {
		d = delay[0]
	}

	return fmt.Sprintf("%s && sleep %f ", dev.SwipeCmd(x, y, x, y), float32(d) / 1000.0)
}
