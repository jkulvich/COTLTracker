package tracker

import (
	"fmt"
	"io"
	"player/tracker/track"
	"player/tracker/track/unit"
	"strings"
	"time"
)

// Stream - Stream tracker.
// It prints minimized COTLTrack file into stream.
// This tracker created as an example and for debugging purposes.
type Stream struct {
	Interface
	// conf - Tracker configuration
	conf StreamConfig
	// stream - Output stream
	stream io.Writer
	// trk - Current track
	trk *track.Track
	// pos - Current playing block pos
	pos int
	// playing - Current inner state of playing
	playing bool
}

// StreamConfig - Stream tracker configuration
type StreamConfig struct {
	// Tick - Length of tick in ms. Default 100
	Tick int
	// Delay - Length of delay between taps. Default 0
	Delay int
}

// NewStream - Create new virtual tracker
func NewStream(stream io.Writer, config StreamConfig) *Stream {
	tracker := Stream{
		stream: stream,
		conf: config,
	}
	return &tracker
}

// Play - Start async playing
func (t *Stream) Play(trk *track.Track) error {
	t.trk = trk
	t.pos = 0
	t.playing = true
	go func() {
		t.resumeLoop()
		t.playing = false
	}()
	return nil
}

// Pause - Pause playing
func (t *Stream) Pause() error {
	t.playing = false
	return nil
}

// Resume - Resume playing
func (t *Stream) Resume() error {
	t.playing = true
	go func() {
		t.resumeLoop()
		t.playing = false
	}()
	return nil
}

// Stop - Stop playing
func (t *Stream) Stop() error {
	t.playing = false
	t.pos = 0
	return nil
}

// State - Return tracker's state
func (t *Stream) State() State {
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
func (t *Stream) SeekBlock(pos int) error {
	t.pos = pos
	return nil
}

// SeekTime - Set cursor position to specific block at time.
// Override this realisation.
func (t *Stream) SeekTime(pos int) error {
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
func (t *Stream) TotalBlocks() int {
	if t.trk != nil {
		return len(t.trk.Units)
	}
	return 0
}

// TotalTime - Length in ms of current track
// Override this realisation.
func (t *Stream) TotalTime() int {
	return t.timeOf(t.trk.Len())
}

// CurrentTime - Current play time
// Override this realisation.
func (t *Stream) CurrentTime() int {
	return t.timeOf(t.pos)
}

// CurrentBlock - Current block position
func (t *Stream) CurrentBlock() int {
	return t.pos
}

// resumeLoop - Start loop of playing
func (t *Stream) resumeLoop() {
	for {
		if !t.playing || t.pos >= len(t.trk.Units) {
			t.playing = false
			break
		}
		// Fetch current unit
		u := t.trk.Units[t.pos]
		// Play note
		if u.Type == unit.TypeNote {
			_, _ = fmt.Fprintf(t.stream, "%s ", u.Note.String())
			time.Sleep(time.Millisecond * time.Duration(t.conf.Delay))
		}
		// Delay
		if u.Type == unit.TypeDelay {
			_, _ = fmt.Fprintf(t.stream, "%s ", strings.Repeat("-", int(u.Delay)))
			time.Sleep(time.Millisecond * time.Duration(t.conf.Tick*int(u.Delay)))
		}
		// Increment cursor position
		t.pos++
		// Exit by end
		if t.State() == StateFinished {
			break
		}
	}
}

// timeOf - Return time in ms for block at pos
func (t *Stream) timeOf(pos int) int {
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