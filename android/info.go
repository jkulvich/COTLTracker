package android

import (
	"errors"
	"regexp"
	"strconv"
)

// ScreenSize - It returns the screen width and height of current device screen
func (dev *Android) ScreenSize() ([2]int, error) {
	res, err := dev.Exec("wm size")
	if err != nil {
		return [2]int{}, err
	}

	parsed := regexp.MustCompile(`(\d+)x(\d+)`).FindStringSubmatch(res)
	if len(parsed) == 3 {
		w, _ := strconv.Atoi(parsed[1])
		h, _ := strconv.Atoi(parsed[2])
		return [2]int{w, h}, nil
	}

	return [2]int{}, errors.New("can't get the device screen size")
}

// ScreenDPI - It returns [physical, override] the device screen
func (dev *Android) ScreenDPI() ([2]int, error) {
	res, err := dev.Exec("wm density")
	if err != nil {
		return [2]int{}, err
	}

	parsed := regexp.MustCompile(`\d+`).FindAllStringSubmatch(res, -1)
	if len(parsed) > 0 {
		physical, _ := strconv.Atoi(parsed[0][0])
		override := 0
		if len(parsed) > 1 {
			override, _ = strconv.Atoi(parsed[1][0])
		}
		if override == 0 {
			override = physical
		}
		return [2]int{physical, override}, nil
	}

	return [2]int{}, errors.New("can't get the device density")
}
