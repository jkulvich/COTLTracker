package parser

import (
	"fmt"
	"player/tracker/track/unit"
	"regexp"
	"strings"
)

// COTLTrack - Works with COTLTrack files, usually with .txt extension
type COTLTrack struct {
	Interface
}

// NewCOTLTrack - Create a new COTLTrack parser
func NewCOTLTrack() *COTLTrack {
	return &COTLTrack{}
}

// Unmarshal - Parse raw track data
func (l *COTLTrack) Unmarshal(data []byte) ([]unit.Unit, error) {
	// Get all lines
	lines := regexp.MustCompile(`([^\n\r]*)`).FindAllString(string(data), -1)

	// Trim whitespaces for every line
	for iLine, line := range lines {
		lines[iLine] = strings.TrimSpace(line)
	}

	// Prepare units' space
	units := make([]unit.Unit, 0, 512)

	// Process every line
	for iLine, line := range lines {
		// Create empty unit
		if len(line) == 0 {
			units = append(units, unit.Unit{})
			continue
		}

		// Comment line
		if strings.HasPrefix(line, "#") {
			u, err := unit.New(line, unit.Extra{Line: iLine})
			if err != nil {
				return nil, err
			}
			units = append(units, u)
			continue
		}

		// Split other units by space
		if !strings.HasPrefix(line, "#") {
			unitsStrings := regexp.MustCompile(`([^\s]+)`).FindAllString(line, -1)

			// Processing every unit and try to create it
			for iUnit, unitStr := range unitsStrings {
				u, err := unit.New(unitStr, unit.Extra{Line: iLine, Pos: iUnit})
				if err != nil {
					return nil, err
				}
				units = append(units, u)
			}
		}
	}

	return units, nil
}

// Marshal - Serialize unit.Unit array to raw track data
func (l *COTLTrack) Marshal(units []unit.Unit) ([]byte, error) {
	return []byte{}, fmt.Errorf("marshaling of cotltrack files not implemented yet")
}
