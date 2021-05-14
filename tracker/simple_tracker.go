package tracker

import (
	"fmt"
	"player/tracker/track"
	"player/tracker/track/unit"
	"time"
)

// SimpleInterface - Another tracker's interface with reduced set of methods.
type SimpleInterface interface {
	// step - Play unit step
	step(data *Simple, n unit.Unit) error
	// beforePlay - Calls each time before play loop activates
	beforePlay(data *Simple) error
	// afterPlay - Calls each time after play loop stops
	afterPlay(data *Simple) error
}

// Simple - Basic realisation of tracker
type Simple struct {
	Interface
	// tracker - Actual managed object
	tracker SimpleInterface
	// trk - Current track
	trk track.Track
	// pos - Current playing block pos
	pos int
	// playing - Current inner state of playing
	playing bool
	// err - Tracker error
	err error
	// Delay - Delay between notes
	Delay int
	// Tick - Single tick count for 'dashes'
	Tick int
}

// Simplify - Takes some work by itself.
// It can be helpful for describing simple trackers which doesn't have complicated functionality.
// So, you can use this to simplify your tracker realisation.
func Simplify(tracker SimpleInterface) *Simple {
	return &Simple{
		tracker: tracker,
	}
}

// Play - Start async playing
func (t *Simple) Play(trk track.Track) {
	t.trk = trk
	t.pos = 0
	t.Resume()
}

// Pause - Pause playing
func (t *Simple) Pause() {
	t.playing = false
}

// Resume - Resume playing
func (t *Simple) Resume() {
	t.playing = true
	if err := t.tracker.beforePlay(t); err != nil {
		t.err = err
		return
	}
	go func() {
		if err := t.resumeLoop(); err != nil {
			t.err = err
			return
		}
		if err := t.tracker.afterPlay(t); err != nil {
			t.err = err
			return
		}
		t.playing = false
	}()
}

// Stop - Stop playing
func (t *Simple) Stop() {
	t.playing = false
	t.pos = 0
}

// State - Return tracker's state
func (t *Simple) State() State {
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
func (t *Simple) SeekBlock(pos int) {
	t.pos = pos
}

// SeekTime - Set cursor position to specific block at time.
// Override this realisation.
func (t *Simple) SeekTime(pos int) {
	var total time.Duration
	for i, u := range t.trk.Units {
		switch u.Type {
		case unit.TypeDelay:
			total += time.Millisecond * time.Duration(u.Delay) * time.Duration(t.Tick)
		case unit.TypeNote:
			total += time.Millisecond * time.Duration(t.Delay)
		}
		if pos < int(total.Milliseconds()) {
			t.pos = i
			return
		}
	}
}

// TotalBlocks - Count of blocks in current track
func (t *Simple) TotalBlocks() int {
	return len(t.trk.Units)
}

// TotalTime - Length in ms of current track
// Override this realisation.
func (t *Simple) TotalTime() int {
	return t.timeOf(t.trk.Len())
}

// CurrentTime - Current play time
// Override this realisation.
func (t *Simple) CurrentTime() int {
	return t.timeOf(t.pos)
}

// CurrentBlock - Current block position
func (t *Simple) CurrentBlock() int {
	return t.pos
}

// Error - Return last error
func (t *Simple) Error(flush ...bool) error {
	err := t.err
	if len(flush) == 1 {
		if flush[0] {
			t.err = nil
		}
	}
	return err
}

// resumeLoop - Start loop of playing
func (t *Simple) resumeLoop() error {
	for {
		if !t.playing || t.pos >= len(t.trk.Units) {
			t.playing = false
			break
		}
		// Play note
		if err := t.tracker.step(t, t.trk.Units[t.pos]); err != nil {
			return err
		}
		// Increment cursor position
		t.pos++
		// Exit by end
		if t.State() == StateFinished {
			break
		}
	}
	return nil
}

// timeOf - Return time in ms for block at pos
func (t *Simple) timeOf(pos int) int {
	var total time.Duration
	for i := 0; i < pos; i++ {
		u := t.trk.Units[i]
		switch u.Type {
		case unit.TypeDelay:
			fmt.Println(">>>>>>>", t.Tick)
			total += time.Millisecond * time.Duration(u.Delay) * time.Duration(t.Tick)
		case unit.TypeNote:
			total += time.Millisecond * time.Duration(t.Delay)
		}
	}
	return int(total.Milliseconds())
}

// step - Unit play step stub
func (t *Simple) step(data *Simple, u unit.Unit) error {
	return nil
}

// beforePlay - Play hook stub
func (t *Simple) beforePlay(data *Simple) error {
	return nil
}

// afterPlay - Play hook stub
func (t *Simple) afterPlay(data *Simple) error {
	return nil
}
