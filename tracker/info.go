package tracker

import (
	"fmt"
	"io"
	"player/tracker/track"
	"player/tracker/track/unit"
	"time"
)

// Info - Info tracker.
// It prints track info like duration, timing, shift and so one into stream.
type Info struct {
	tracker
	// conf - Tracker configuration
	conf InfoConfig
	// stream - Output stream
	stream io.Writer
}

// InfoConfig - Info tracker configuration
type InfoConfig struct {
	// Tick - Length of tick in ms. Default 100
	Tick int
	// Delay - Length of delay between taps. Default 0
	Delay int
}

// NewInfo - Create new virtual tracker
func NewInfo(stream io.Writer, config ...InfoConfig) *Info {
	tracker := Info{
		stream: stream,
		conf: InfoConfig{
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
func (t *Info) resumeLoop() {
	totalTime := time.Duration(t.TotalTime()) * time.Millisecond
	_, _ = fmt.Fprintf(t.stream, "Duration: %s (%dms)\n", totalTime.String(), totalTime.Milliseconds())
	_, _ = fmt.Fprintf(t.stream, "Blocks: %d\n", t.trk.Len())
	_, _ = fmt.Fprintf(t.stream, "Timing: %d\n", t.trk.GetTiming())
	_, _ = fmt.Fprintf(t.stream, "Shift: %d\n", t.trk.GetShift())
}

// timeOf - Return time in ms for block at pos
func (t *Info) timeOf(pos int) int {
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
func (t *Info) SeekTime(pos int) error {
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
func (t *Info) TotalTime() int {
	return t.timeOf(t.trk.Len())
}

// CurrentTime - Current play time
// Override this realisation.
func (t *Info) CurrentTime() int {
	return t.timeOf(t.pos)
}

// Play - Start async playing
func (t *Info) Play(trk *track.Track) error {
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
func (t *Info) Resume() error {
	t.playing = true
	go func() {
		t.resumeLoop()
		t.playing = false
	}()
	return nil
}