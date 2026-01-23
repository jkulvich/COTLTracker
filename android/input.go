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

	cmd := fmt.Sprintf("input swipe %d %d %d %d %d", x1, y1, x2, y2, d)
	if _, err := dev.Exec(cmd); err != nil {
		return err
	}
	return nil
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

func (dev *Android) MultiTouchTap(x, y int, delay ...int) {
	d := dev.MinTapDelay
	if len(delay) > 0 {
		d = delay[0]
	}

	go func() {
		cmd := fmt.Sprintf("sendevent /dev/input/eventX 3 57 0 && sendevent /dev/input/eventX 3 53 %d && sendevent /dev/input/eventX 3 54 %d && sendevent /dev/input/eventX 0 2 0 && sendevent /dev/input/eventX 3 57 -1 && /dev/input/eventX 0 2 0", x, y)
		res, err := dev.Exec(cmd)
		fmt.Println(res, err)
	}()
	<-time.After(time.Millisecond * time.Duration(d))
}
