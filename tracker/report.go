package tracker

import (
	"fmt"
	"gopkg.in/gookit/color.v1"
	"player/tracker/track/unit"
	"time"
)

// Report - Report tracker.
// It prints track info like duration, timing, shift and so one into stream.
type Report struct {
	Simple
}

// ReportConfig - Configuration of Report tracker
type ReportConfig struct {
	// Tick - Length of tick in ms for delay nodes like 'dashes'
	Tick int
	// Delay - Length of delay between taps to prevent taps throttling
	Delay int
}

// NewReport - Create new report tracker
func NewReport(config ReportConfig) *Simple {
	tracker := &Report{}
	tracker.Delay = config.Delay
	tracker.Tick = config.Tick
	return Simplify(tracker)
}

// beforePlay - Unit play step
func (t *Report) beforePlay(data *Simple) error {
	c := struct {
		Title color.Color
		Param color.Color
		Good  color.Color
		Warn  color.Color
		Bad   color.Color
	}{
		Title: color.LightCyan,
		Param: color.LightWhite,
		Good:  color.Green,
		Warn:  color.Yellow,
		Bad:   color.Red,
	}

	c.Title.Printf("===== COTLTrack Report =====")
	fmt.Println()

	// Duration
	totalTime := time.Duration(t.TotalTime()) * time.Millisecond
	c.Param.Printf("%-15s: ", "Duration")
	if totalTime <= time.Minute*3 {
		c.Good.Printf("%s or %dms | Normal duration", totalTime.String(), totalTime.Milliseconds())
	} else {
		c.Warn.Printf("%s or %dms | Too long track", totalTime.String(), totalTime.Milliseconds())
		c.Warn.Printf("\n\t%s", "Your track's duration too long.")
		c.Warn.Printf("\n\t%s", "It can be hard to Simplify this track.")
	}
	fmt.Println()

	// Blocks
	c.Param.Printf("%-15s: ", "Total Blocks")
	if t.trk.Len() <= int(totalTime.Milliseconds())/120 {
		c.Good.Printf("%d | Blocks count match with duration", t.trk.Len())
	} else {
		c.Warn.Printf("%d | Too many blocks for current duration", t.trk.Len())
		c.Warn.Printf("\n\t%s", "Unusual count of blocks detected.")
		c.Warn.Printf("\n\t%s", "Too many comments, empty lines or separated dashes?")
	}
	fmt.Println()

	// Timing & tick duration
	c.Param.Printf("%-15s: ", "Timing")
	timing := t.trk.GetTiming()
	if timing >= 100 && timing <= 200 {
		c.Good.Printf("%dms | Delay is in normal range", timing)
	} else {
		c.Warn.Printf("%dms | Unusual delay value", timing)
		c.Warn.Printf("\n\t%s", "Usually, value should be between 100ms and 200ms.")
		c.Warn.Printf("\n\t%s", "Your 'dash' delay is too short or long.")
		c.Warn.Printf("\n\t%s", "It can cause troubles if you want Simplify your track more accuracy.")
	}
	fmt.Println()

	// Timing override
	c.Param.Printf("%-15s: ", "Default Timing")
	if timing == t.Tick {
		c.Good.Printf("YES | Track plays as expected")
	} else {
		c.Warn.Printf("NO | Track plays faster or slower", timing)
		c.Warn.Printf("\n\t%s", "Track's timing was overridden by the app.")
		c.Warn.Printf("\n\t%s", "It can be helpful to make your track slower or faster.")
		c.Warn.Printf("\n\t%s", "However, it can cause unexpected sounds.")
	}
	fmt.Println()

	// Shift info
	c.Param.Printf("%-15s: ", "Shift Value")
	shift := t.trk.GetShift()
	if shift == 0 {
		c.Good.Printf("%d | Normalized", shift)
	} else {
		c.Warn.Printf("%d | Isn't normalized", shift)
		c.Warn.Printf("\n\t%s", "Track isn't normalized and it can cause unexpected behaviour.")
		c.Warn.Printf("\n\t%s", "In usual way your track should be normalized to set first")
		c.Warn.Printf("\n\t%s", "track's octave equal N0, where N is note char.")
		c.Warn.Printf("\n\t%s", "Manual octave shifting can be helpful in very rare cases.")
	}
	fmt.Println()

	// Delay info
	c.Param.Printf("%-15s: ", "Tap Delay")
	delay := t.Delay
	if delay >= 40 && delay <= 100 {
		c.Good.Printf("%d | Tap delay is in normal range", delay)
	} else {
		c.Warn.Printf("%d | Unusual tap delay", delay)
		c.Warn.Printf("\n\t%s", "Usual value is in range 40ms to 100ms for high-latency devices.")
		c.Warn.Printf("\n\t%s", "Value lower than 40ms can cause 'tap' throttling.")
		c.Warn.Printf("\n\t%s", "Value greater than 100ms better to replace with 'dashing'.")
		c.Warn.Printf("\n\t%s", "So you can Simplify your song more accuracy.")
	}
	fmt.Println()

	// Check for 2.5 octaves
	c.Param.Printf("%-15s: ", "Overflow")
	trk := t.trk.Clone()
	trk.Normalize()
	octavesOverflow := -1
	maxNote, _ := unit.NewNote("C2")
	for i, u := range trk.Units {
		if u.Type == unit.TypeNote && u.Note > maxNote {
			octavesOverflow = i
			break
		}
	}
	if octavesOverflow == -1 {
		c.Good.Printf("NO | Track can plays with CotL")
	} else {
		c.Warn.Printf("YES | At position %d", octavesOverflow)
		c.Warn.Printf("\n\t%s", "Octaves overflow detected.")
		c.Warn.Printf("\n\t%s", "It means, that this track can't be played normal")
		c.Warn.Printf("\n\t%s", "in Sky: Children of the Light.")
		c.Warn.Printf("\n\t%s", "Decrease the note to fix.")
	}
	fmt.Println()

	t.playing = false
	t.pos = len(t.trk.Units)
	return nil
}
