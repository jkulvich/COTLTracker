package parser

import (
	"player/tracker/track/unit"
)

// Interface - Interface for track parsers.
// Parsers can marshal and unmarshal raw track data
type Interface interface {
	// Unmarshal raw track data to unit.Unit array
	Unmarshal(data []byte) ([]unit.Unit, error)
	// Marshal unit.Unit array to raw track data
	Marshal(units []unit.Unit) ([]byte, error)
}
