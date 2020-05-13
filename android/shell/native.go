package shell

import (
	"os/exec"
	"strings"
)

// Native - The native shell for current system like sh
type Native struct {
	shell string
	args  []string
}

// NewNative - It creates a new native shell instance to run command under current machine
func NewNative() (*Native, error) {
	return &Native{
		shell: "sh",
		args:  []string{"-c"},
	}, nil
}

// Exec - It runs command under current shell
func (sh Native) Exec(command string) (string, error) {
	args := append(sh.args, command)
	cmd := exec.Command(sh.shell, args...)
	resp, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(resp)), nil
}

// Type - It returns shell name
func (sh Native) Type() string {
	return "native"
}
