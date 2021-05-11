// track - Module to parse, convert and manipulate tracks

package track

import (
	"math"
	"player/tracker/track/unit"
)

// Track - Musical file
type Track struct {
	// Units - Array of track file units like notes, delays, comments etc...
	Units []unit.Unit
}

// New - Create new empty track
func New() *Track {
	return &Track{}
}

// GetTiming - Get track timing.
// It tries find timing comment and return value or 0 if not found.
func (t *Track) GetTiming() int {
	for _, u := range t.Units {
		if u.Type == unit.TypeComment && u.Comment.Tag() == unit.CommentTagTiming {
			return u.Comment.ValueInt()
		}
	}
	return 0
}

// GetShift - Get track shift.
// It tries find shift comment and return value or 0 if not found.
func (t *Track) GetShift() int {
	for _, u := range t.Units {
		if u.Type == unit.TypeComment && u.Comment.Tag() == unit.CommentTagShift {
			return u.Comment.ValueInt()
		}
	}
	return 0
}

// Shift - Shift all notes.
func (t *Track) Shift(shift int) {
	for i := range t.Units {
		if t.Units[i].Type == unit.TypeNote {
			t.Units[i].Note += unit.Note(shift)
		}
	}
}

// Normalize - Normalize all notes by 1 octave.
// Normalized notes - range [C0, C2]
func (t *Track) Normalize() {
	// Find minimal octave
	minOct := math.MaxInt32
	for _, u := range t.Units {
		if u.Type == unit.TypeNote && u.Note.Octave() < minOct {
			minOct = u.Note.Octave()
		}
	}
	// If value didn't change
	if minOct == math.MaxInt32 {
		minOct = 0
	}

	// Shift to normalize
	t.Shift(-minOct * 7)
}

// Len - Count of units
func (t *Track) Len() int {
	return len(t.Units)
}

// Sub - Slice new track from original
func (t *Track) Sub(start, end int) *Track {
	return &Track{
		Units: t.Units[start:end],
	}
}
