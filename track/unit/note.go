package unit

import (
	"fmt"
	"regexp"
	"strconv"
)

// Note - note in hash number
type Note uint8

// NewNote - Parse note from letter notation string and return note in a hash number.
// Valid letter value from C0 to F36.
// If you need to shift note tone just add value to this variable like: n += 1.
// If you need to shift octave add 7 like: n += 7.
func NewNote(str string) (Note, error) {
	// Checking for note format
	reg := regexp.MustCompile(`(?i)^([A-H])(\d*)$`)
	if !reg.MatchString(str) {
		return 0, fmt.Errorf("incorrect note letter format: %s", str)
	}
	// Fetch note components & replace "H" to "B" for more easy processing
	tone := str[0]
	octave, _ := strconv.Atoi(str[1:])
	if tone == 'H' {
		tone = 'B'
	}
	// Replacement map to make correct shifting support
	rpl := map[uint8]uint8{'C': 'A', 'D': 'B', 'E': 'C', 'F': 'D', 'G': 'E', 'A': 'F', 'B': 'G'}
	hash := int(rpl[tone]-'A') + octave*7
	if hash < 0 || hash > 255 {
		return 0, fmt.Errorf("out of boundary [C0-F36]: %s", str)
	}
	return Note(hash), nil
}

// String - Note in letter notation like A2, B4, C6, etc...
func (n Note) String() string {
	return fmt.Sprintf("%c%d", n.Tone(), n.Octave())
}

// Tone - Return note's tone letter like A, B, C, D, E, F, G.
// If the note has "H" tone, "B" will be returned instead for compatibility.
func (n Note) Tone() uint8 {
	// Back replacement map to restore actual tone letter
	rplBack := map[uint8]uint8{'A': 'C', 'B': 'D', 'C': 'E', 'D': 'F', 'E': 'G', 'F': 'A', 'G': 'B'}
	return rplBack[uint8(int('A')+int(n)%7)]
}

// Octave - Return note's octave number like 1, 2, 3, etc...
func (n Note) Octave() int {
	return int(n / 7)
}
