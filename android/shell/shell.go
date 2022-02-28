package shell

import (
	"regexp"
)

// Sheller - Universal interface to work with different shells
type Sheller interface {
	Exec(string) (string, error)
	Type() string
}

// New - It detects and creates universal shell for PC+Android or Android only shell
func New() (Sheller, error) {
	native, err := NewNative()
	if err != nil {
		return nil, err
	}

	android, err := isAndroid(native)
	if err != nil {
		return nil, err
	}

	if !android {
		sh, err := NewADB(native)
		if err != nil {
			return nil, err
		}
		return sh, nil
	}

	return native, nil
}

// isAndroid - It checks current OS is Android
func isAndroid(sh *Native) (bool, error) {
	resp, err := sh.Exec("uname")
	if err != nil {
		return false, err
	}

	return regexp.MustCompile(`(?i)android`).MatchString(resp), nil
}
