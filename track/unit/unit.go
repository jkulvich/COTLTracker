package unit

import "fmt"

// Type - Describes type of unit block.
// It can be Note, Delay, Comment or something else
type Type int

const (
	// TypeEmpty - Empty type for just created Units
	TypeEmpty Type = iota
	// TypeNote - Note type
	TypeNote Type = iota
	// TypeDelay - Delay type
	TypeDelay Type = iota
	// TypeComment - Comment type
	TypeComment Type = iota
)

// Unit - A smallest part of track which describes note or delay or something else
type Unit struct {
	// Type - Unit type
	Type Type
	// Extra - Some additional and debug info about Unit
	Extra Extra
	// Delay in ticks (not in ms as early). Now only track player decides time of tick
	Delay Delay
	// Note as a hash number
	Note Note
	// Comment string
	Comment Comment
}

// NewUnit - Make new unit and detect type from string
func NewUnit(str string, extra ...Extra) (Unit, error) {
	var ext Extra
	if len(extra) == 1 {
		ext = extra[0]
	}

	// Detecting note
	if note, err := NewNote(str); err == nil {
		return Unit{
			Type: TypeNote,
			Extra: ext,
			Note: note,
		}, nil
	}

	// Detecting delay
	if delay, err := NewDelay(str); err == nil {
		return Unit{
			Type: TypeDelay,
			Extra: ext,
			Delay: delay,
		}, nil
	}

	// Detecting comment
	if comment, err := NewComment(str); err == nil {
		return Unit{
			Type: TypeComment,
			Extra: ext,
			Comment: comment,
		}, nil
	}

	// Print error info with extra info
	if ext.Pos != 0 || ext.Line != 0 {
		return Unit{}, fmt.Errorf("can't detect any known type at pos %d:%d: %s", ext.Line, ext.Pos, str)
	}
	return Unit{}, fmt.Errorf("can't detect any known type: %s", str)
}
