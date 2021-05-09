// track - Module to parse, convert and manipulate tracks

package track

import "player/track/unit"

// Track - Musical file
type Track struct {
	// Units - Array of track file units like notes, delays, comments etc...
	Units []unit.Unit
}
