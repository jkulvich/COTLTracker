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
	for i := t.pos; i < len(t.trk.Units); i++ {
		if !t.playing {
			break
		}
		// Fetch current unit
		u := t.trk.Units[i]
		// Play note
		if u.Type == unit.TypeNote {
			fmt.Printf("%s ", u.Note.String())
			time.Sleep(time.Millisecond * time.Duration(t.conf.Delay))
		}
		// Delay
		if u.Type == unit.TypeDelay {
			fmt.Printf("%s ", strings.Repeat("-", int(u.Delay)))
			time.Sleep(time.Millisecond * time.Duration(t.conf.Tick*int(u.Delay)))
		}
	}
}

// SeekTime - Set cursor position to specific block.
// Override this realisation.
func (t *Stream) SeekTime(pos int) error {
	panic("don't use pure realisation of seek time in tracker")
}

// TotalTime - Length in ms of current track
// Override this realisation.
func (t *Stream) TotalTime() int {
	panic("don't use pure realisation of total time in tracker")
}

// CurrentTime - Current play time
// Override this realisation.
func (t *Stream) CurrentTime() int {
	panic("don't use pure realisation of current time in tracker")
}
