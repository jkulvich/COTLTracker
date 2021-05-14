package tracker

import (
	"fmt"
	"io"
	"player/tracker/track/unit"
	"strings"
	"time"
)

// Stream - Stream tracker.
// It prints minimized COTLTrack file into stream.
// This tracker created as an example and for debugging purposes.
type Stream struct {
	Simple
	// stream - Output stream
	stream io.Writer
}

// StreamConfig - Configuration of Stream tracker
type StreamConfig struct {
	// Tick - Length of tick in ms for delay nodes like 'dashes'
	Tick int
	// Delay - Length of delay between taps to prevent taps throttling
	Delay int
}

// NewStream - Create new virtual tracker
func NewStream(stream io.Writer, config StreamConfig) *Simple {
	tracker := &Stream{
		stream: stream,
	}
	tracker.Delay = config.Delay
	tracker.Tick = config.Tick
	return Simplify(tracker)
}

// step - Unit play step
func (t *Stream) step(data *Simple, u unit.Unit) error {
	if u.Type == unit.TypeNote {
		_, _ = fmt.Fprintf(t.stream, "%s ", u.Note.String())
		time.Sleep(time.Millisecond * time.Duration(t.Delay))
	}
	// Delay
	if u.Type == unit.TypeDelay {
		_, _ = fmt.Fprintf(t.stream, "%s ", strings.Repeat("-", int(u.Delay)))
		time.Sleep(time.Millisecond * time.Duration(t.Tick*int(u.Delay)))
	}
	return nil
}
