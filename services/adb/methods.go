package adb

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// GetVendorModel - Return vendor and model of device
func (s *ADBService) GetVendorModel() (string, error) {
	// Get device vendor & model
	resp, err := s.dev.RunCommand("getprop ro.product.vendor.brand && getprop ro.product.model")
	if err != nil {
		return "", err
	}
	// Remove \n's
	venmodel := strings.ReplaceAll(resp, "\n", " ")
	return venmodel, nil
}

// GetScreenSize - Return actual screen size in pixels
func (s *ADBService) GetScreenSize() ([2]int, error) {
	// Get device vendor & model
	resp, err := s.dev.RunCommand("wm size")
	if err != nil {
		return [2]int{}, err
	}

	// Parse result
	sizes := regexp.MustCompile(`(\d+)x(\d+)`).FindStringSubmatch(resp)
	if len(sizes) != 3 {
		return [2]int{}, errors.New(fmt.Sprintf("result contains something strange: %s", resp))
	}

	w, _ := strconv.Atoi(sizes[1])
	h, _ := strconv.Atoi(sizes[2])

	return [2]int{w, h}, nil
}

// GetScreenSizeAlbum - Return actual screen size in album mode
func (s *ADBService) GetScreenSizeAlbum() ([2]int, error) {
	sizes, err := s.GetScreenSize()
	if err != nil {
		return [2]int{}, err
	}

	if sizes[0] > sizes[1] {
		return [2]int{sizes[0], sizes[1]}, nil
	} else {
		return [2]int{sizes[1], sizes[0]}, nil
	}
}

// Tap - Make fast tap on screen
func (s *ADBService) Tap(x, y int) error {
	const safeDelay = 40
	_, err := s.dev.RunCommand(fmt.Sprintf("input swipe %[1]d %[2]d %[1]d %[2]d %[3]d", x, y, safeDelay))
	if err != nil {
		return err
	}
	return nil
}
