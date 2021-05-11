package tracker

import (
	"fmt"
	"gopkg.in/gookit/color.v1"
	"player/tracker/track"
	"player/tracker/track/unit"
	"time"
)

// Report - Report tracker.
// It prints track info like duration, timing, shift and so one into stream.
type Report struct {
	Interface
	// conf - Tracker configuration
	conf ReportConfig
	// trk - Current track
	trk *track.Track
	// pos - Current playing block pos
	pos int
	// playing - Current inner state of playing
	playing bool
}

// ReportConfig - Report tracker configuration
type ReportConfig struct {
	// Tick - Length of tick in ms. Default 100
	Tick int
	// Delay - Length of delay between taps. Default 0
	Delay int
}

// NewReport - Create new report tracker
func NewReport(config ReportConfig) *Report {
	tracker := Report{
		conf: config,
	}
	return &tracker
}

// Play - Start async playing
func (t *Report) Play(trk *track.Track) error {
	t.trk = trk
	go func() {
		t.resumeLoop()
		t.playing = false
	}()
	return nil
}

// Pause - Pause playing
func (t *Report) Pause() error {
	t.playing = false
	return nil
}

// Resume - Resume playing
func (t *Report) Resume() error {
	t.playing = true
	go func() {
		t.resumeLoop()
		t.playing = false
	}()
	return nil
}

// Stop - Stop playing
func (t *Report) Stop() error {
	t.playing = false
	t.pos = 0
	return nil
}

// State - Return tracker's state
func (t *Report) State() State {
	if t.playing {
		return StatePlaying
	}
	if t.pos == 0 {
		return StateStopped
	}
	if t.pos >= len(t.trk.Units) {
		return StateFinished
	}
	return StatePaused
}

// SeekBlock - Set cursor position to specific block
func (t *Report) SeekBlock(pos int) error {
	t.pos = pos
	return nil
}

// SeekTime - Set cursor position to specific block at time.
// Override this realisation.
func (t *Report) SeekTime(pos int) error {
	var total time.Duration
	for i, u := range t.trk.Units {
		switch u.Type {
		case unit.TypeDelay:
			total += time.Millisecond * time.Duration(u.Delay) * time.Duration(t.conf.Tick)
		case unit.TypeNote:
			total += time.Millisecond * time.Duration(t.conf.Delay)
		}
		if pos < int(total.Milliseconds()) {
			t.pos = i
			return nil
		}
	}
	return nil
}

// TotalBlocks - Count of blocks in current track
func (t *Report) TotalBlocks() int {
	if t.trk != nil {
		return len(t.trk.Units)
	}
	return 0
}

// TotalTime - Length in ms of current track
// Override this realisation.
func (t *Report) TotalTime() int {
	return t.timeOf(t.trk.Len())
}

// CurrentTime - Current play time
// Override this realisation.
func (t *Report) CurrentTime() int {
	return t.timeOf(t.pos)
}

// CurrentBlock - Current block position
func (t *Report) CurrentBlock() int {
	return t.pos
}

// resumeLoop - Start loop of playing
func (t *Report) resumeLoop() {
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
		c.Warn.Printf("\n\t%s", "It can be hard to manage this track.")
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
		c.Warn.Printf("\n\t%s", "It can cause troubles if you want manage your track more accuracy.")
	}
	fmt.Println()

	// Timing override
	c.Param.Printf("%-15s: ", "Default Timing")
	if timing == t.conf.Tick {
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
	delay := t.conf.Delay
	if delay >= 40 && delay <= 100 {
		c.Good.Printf("%d | Tap delay is in normal range", delay)
	} else {
		c.Warn.Printf("%d | Unusual tap delay", delay)
		c.Warn.Printf("\n\t%s", "Usual value is in range 40ms to 100ms for high-latency devices.")
		c.Warn.Printf("\n\t%s", "Value lower than 40ms can cause 'tap' throttling.")
		c.Warn.Printf("\n\t%s", "Value greater than 100ms better to replace with 'dashing'.")
		c.Warn.Printf("\n\t%s", "So you can manage your song more accuracy.")
	}
	fmt.Println()

	t.pos = len(t.trk.Units)
	t.playing = false
}

// timeOf - Return time in ms for block at pos
func (t *Report) timeOf(pos int) int {
	var total time.Duration
	for i := 0; i < pos; i++ {
		u := t.trk.Units[i]
		switch u.Type {
		case unit.TypeDelay:
			total += time.Millisecond * time.Duration(u.Delay) * time.Duration(t.conf.Tick)
		case unit.TypeNote:
			total += time.Millisecond * time.Duration(t.conf.Delay)
		}
	}
	return int(total.Milliseconds())
}
