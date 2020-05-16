package controller

import (
	"fmt"
	"player/android"
)

// Controller - The structure to control "Children Of The Light" gameplay
type Controller struct {
	dev    *android.Android
	screen [2]int
}

// New - It creates a new controller instance
func New(dev *android.Android) (*Controller, error) {
	screen, err := dev.ScreenSize()
	if err != nil {
		return nil, err
	}

	// If incorrect screen layout
	if screen[1] > screen[0] {
		screen = [2]int{screen[1], screen[0]}
	}

	return &Controller{
		dev:    dev,
		screen: screen,
	}, nil
}

// harpButtonPos - It calculates screen position of harp button by button position
func (ctrl *Controller) harpButtonPos(x, y int) ([2]int, error) {
	if x < 0 || y < 0 || x > 4 || y > 2 {
		return [2]int{}, fmt.Errorf("incorrect btn position, got [%d, %d], must be inside [0-4, 0-2]", x, y)
	}

	// Screen size
	scrWidth := float64(ctrl.screen[0])
	scrHeight := float64(ctrl.screen[1])

	// Harp panel size
	panelHeight := scrHeight * .57   //< Constant for Harp panel height
	panelWidth := panelHeight * 1.65 //< Constant for panel width

	// Harp panel buttons sizes
	btnWidth := panelWidth / 5.
	btnHeight := panelHeight / 3.

	// Harp panel margins
	marginLeft := (scrWidth - panelWidth) / 2.
	marginTop := scrWidth * .04 //< Constant for Harp panel margin

	// Calc screen coordinates
	pointX := int(marginLeft + btnWidth*float64(x) + btnWidth/2.)
	pointY := int(marginTop + btnHeight*float64(y) + btnHeight/2.)

	return [2]int{pointX, pointY}, nil
}

// HarpTap - It taps on harp button
func (ctrl *Controller) HarpTap(x, y int) error {
	pos, err := ctrl.harpButtonPos(x, y)
	if err != nil {
		return err
	}

	ctrl.dev.FastTap(pos[0], pos[1])
	return nil
}

type tone map[uint8][]int
type octave map[int]tone

// HarpTapNote - It taps on harp note
func (ctrl *Controller) HarpTapNote(oct int, ton byte) error {
	var octaves = octave{
		0: tone{
			'C': []int{0, 0},
			'D': []int{1, 0},
			'E': []int{2, 0},
			'F': []int{3, 0},
			'G': []int{4, 0},
			'A': []int{0, 1},
			'B': []int{1, 1},
		},
		1: tone{
			'C': []int{2, 1},
			'D': []int{3, 1},
			'E': []int{4, 1},
			'F': []int{0, 2},
			'G': []int{1, 2},
			'A': []int{2, 2},
			'B': []int{3, 2},
		},
		2: tone{
			'C': []int{4, 2},
		},
	}
	btn := octaves[oct][ton]
	if len(btn) != 2 {
		return fmt.Errorf("incorrect note tone or octave for play %c%d, use A0-D1", ton, oct)
	}
	return ctrl.HarpTap(btn[0], btn[1])
}
