package parser

import (
	"fmt"
	"player/track/unit"
)

// Interface - Interface for track parsers.
// Parsers can marshal and unmarshal raw track data
type Interface interface {
	// Unmarshal raw track data to unit.Unit array
	Unmarshal(data []byte) ([]unit.Unit, error)
	// Marshal unit.Unit array to raw track data
	Marshal(units []unit.Unit) ([]byte, error)
}

// Parser - Basic realisation of parser without data parsing
type Parser struct{}

// Unmarshal - Unmarshal stub
func (l *Parser) Unmarshal(data []byte) ([]unit.Unit, error) {
	return []unit.Unit{}, fmt.Errorf("don't use this parser, this is just an abstract class")
}

// Marshal - Marshal stub
func (l *Parser) Marshal(units []unit.Unit) ([]byte, error) {
	return []byte{}, fmt.Errorf("don't use this parser, this is just an abstract class")
}
