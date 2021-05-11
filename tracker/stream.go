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
	tracker
	// conf - Tracker configuration
	conf StreamConfig
	// stream - Output stream
	stream io.Writer
}

// StreamConfig - Stream tracker configuration
type StreamConfig struct {
	// Tick - Length of tick in ms. Default 100
	Tick int
	// Delay - Length of delay between taps. Default 0
	Delay int
}

// NewStream - Create new virtual tracker
func NewStream(stream io.Writer, config ...StreamConfig) *Stream {
	tracker := Stream{
		stream: stream,
		conf: StreamConfig{
			Tick:  0,
			Delay: 0,
		},
	}

	// Override default config
	if len(config) == 1 {
		c := config[0]
		if c.Tick != 0 {
			tracker.conf.Tick = c.Tick
		}
		if c.Delay != 0 {
			tracker.conf.Delay = c.Delay
		}
	}

	return &tracker
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

// Resume - Resume playing
func (t *Stream) Resume() error {
	t.playing = true
	go func() {
		t.resumeLoop()
		t.playing = false
	}()
	return nil
}