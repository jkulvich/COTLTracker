// device - Service module to communicate with ADB

package device

import (
	"errors"
	"fmt"
	"github.com/zach-klippenstein/goadb"
)

// Device - Service to communicate with ADB
type Device struct {
	client *adb.Adb
	dev    *adb.Device
}

// NewDevice - Create new instance of ADB service
func NewDevice() (*Device, error) {
	// Find & start ADB client
	client, err := initADB()
	if err != nil {
		return nil, errors.New(fmt.Sprintf("can't init device: %s", err))
	}

	// Only 1 device must be connected
	serials, _ := client.ListDeviceSerials()
	if len(serials) != 1 {
		return nil, errors.New(fmt.Sprintf("only one device must be connected, found: %d", len(serials)))
	}

	// Connect to any available device
	dev := client.Device(adb.AnyDevice())
	state, err := dev.State()
	if err != nil {
		return nil, errors.New(fmt.Sprintf("can't init device: %s", err))
	}
	if state != adb.StateOnline {
		return nil, errors.New(fmt.Sprintf("can't init device: device is not online"))
	}

	return &Device{
		client: client,
		dev:    dev,
	}, nil
}

// initADB - Find ADB near this app or use ADB from PATH var to start it & return client instance
func initADB() (*adb.Adb, error) {
	// Package for Windows: https://dl.google.com/android/repository/platform-tools-latest-windows.zip
	// How to install ADB: https://www.xda-developers.com/install-adb-windows-macos-linux/

	// Possible ADB locations
	adbPaths := []string{
		"device",
		"C:/platform-tools/device",
		"C:/ADB/device",
	}

	// Trying to find device client in PATH
	client, err := adb.New()
	if err != nil {
		// Trying to find device client in possible paths
		for _, adbPath := range adbPaths {
			client, err = adb.NewWithConfig(adb.ServerConfig{
				PathToAdb: adbPath,
			})
			if err == nil {
				break
			}
		}
	}

	// If ADB found - start it
	if client != nil {
		err = client.StartServer()
	}

	return client, err
}
