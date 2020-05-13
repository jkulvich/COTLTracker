package android

import (
	"player/android/shell"
)

type Android struct {
	shell       shell.Sheller
	MinTapDelay int
}

// New - It creates a new Android device control instance.
// ADB or Native shell will selected automatically.
// It depends on device where you're running on.
func New() (*Android, error) {
	sh, err := shell.New()
	if err != nil {
		return nil, err
	}

	return &Android{
		shell:       sh,
		MinTapDelay: 200,
	}, nil
}

// ShellName - It returns current shell which in use now. Native or ADB
func (dev *Android) ShellName() string {
	return dev.shell.Type()
}

// Exec - It runs a command
func (dev *Android) Exec(command string) (string, error) {
	return dev.shell.Exec(command)
}
