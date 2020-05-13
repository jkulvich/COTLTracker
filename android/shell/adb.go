package shell

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path"
	"regexp"
)

// ADB - ADB Shell implementation.
// Actually, it's decorator for Native Shell
type ADB struct {
	shell  *Native
	adb    string
	serial string
}

// NewADB - It creates a new ADB shell
func NewADB(shell *Native) (*ADB, error) {
	adb, err := findADB()
	if err != nil {
		return nil, err
	}

	serials, err := getDevices(shell)
	if err != nil {
		return nil, err
	}
	if len(serials) == 0 {
		return nil, errors.New("can't find any authorized adb device")
	}
	if len(serials) > 1 {
		return nil, errors.New("too many authorized devices must be only one")
	}
	serial := serials[0]

	return &ADB{
		shell:  shell,
		adb:    adb,
		serial: serial,
	}, nil
}

// Exec - It runs command on connected ADB device
func (adb ADB) Exec(command string) (string, error) {
	cmd := fmt.Sprintf("%s -s %s shell %s", adb.adb, adb.serial, command)
	return adb.shell.Exec(cmd)
}

// Type - It returns shell name
func (adb ADB) Type() string {
	return "adb"
}

// getDevices - It returns list of serials
func getDevices(shell *Native) ([]string, error) {
	res, err := shell.Exec("adb devices")
	if err != nil {
		return nil, err
	}

	parsed := regexp.MustCompile(`(?i)([\dabcdef]{8})\s+(.+)`).FindAllStringSubmatch(res, -1)

	devs := make([]string, 0, len(parsed))
	for i := range parsed {
		status := parsed[i][2]
		if status != "unauthorized" {
			devs = append(devs, parsed[i][1])
		}
	}

	return devs, nil
}

// findADB - It finds ADB in $PATH or local and returns path
func findADB() (string, error) {
	aliases := []string{"./adb", "adb"}

	for _, alias := range aliases {
		if alias == path.Base(alias) {

			// Case when the app should be in $PATH
			p, err := exec.LookPath(alias)
			if err != nil {
				if errors.Is(exec.ErrNotFound, err) {
					continue
				}
				return "", err
			}
			return p, nil

		} else {

			// Case when the app should be near this app
			if _, err := os.Stat(alias); os.IsNotExist(err) {
				continue
			}
			return alias, nil

		}
	}

	return "", fmt.Errorf("can't find adb")
}
