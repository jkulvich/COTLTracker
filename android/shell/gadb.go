package shell

import (
	"fmt"

	"github.com/electricbubble/gadb"
)

// ADB - ADB Shell implementation.
// Actually, it's decorator for Native Shell
type GADB struct {
	dev gadb.Device
}

// NewADB - It creates a new embedded ADB shell
func NewGADB(_ *Native) (*GADB, error) {
	adb, err := gadb.NewClient()
	if err != nil {
		return nil, fmt.Errorf("unable to create gadb client: %w", err)
	}

	devices, err := adb.DeviceList()
	if err != nil {
		return nil, fmt.Errorf("unable to get list of devices: %w", err)
	}

	if len(devices) == 0 {
		return nil, fmt.Errorf("can't find any authorized adb device")
	}
	if len(devices) > 1 {
		return nil, fmt.Errorf("too many authorized devices must be only one")
	}

	device := devices[0]

	return &GADB{
		dev: device,
	}, nil
}

// Exec - It runs command on connected ADB device
func (adb GADB) Exec(command string) (string, error) {
	return adb.dev.RunShellCommand(command)
}

// Type - It returns shell name
func (adb GADB) Type() string {
	return "gadb"
}
