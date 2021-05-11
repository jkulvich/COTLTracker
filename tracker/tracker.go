package tracker

import (
	"player/tracker/track"
)

// State - Tracker play state
type State int

const (
	// StatePaused - Track in pause now
	StatePaused State = iota
	// StatePlaying - Track is playing now
	StatePlaying State = iota
	// StateStopped - Same as StatePause but cursor at zero position
	StateStopped State = iota
	// StateFinished - Same as StatePause but cursor at last position
	StateFinished State = iota
)

// Interface - Trackers' interface.
// Tracker using to manage track playing.
// Play and Resume should be async functions and shouldn't block current goroutine
type Interface interface {
	// Play - Start async play
	// Need to be overridden
	Play(trk *track.Track) error
	// Pause - Pause playing
	Pause() error
	// Resume - Resume async play
	// Need to be overridden
	Resume() error
	// Stop - Pause and set cursor position to start
	Stop() error
	// State - Tracker's current state like StateStopped, StatePlaying, StatePaused and StateFinished
	State() State
	// SeekBlock - Set cursor position to specific block
	SeekBlock(pos int) error
	// SeekTime - Set cursor position to specific block at track time.
	// Need to be overridden
	SeekTime(ms int) error
	// TotalBlocks - Count of blocks in current track
	TotalBlocks() int
	// TotalTime - Length in ms of current track
	// Need to be overridden
	TotalTime() int
	// CurrentBlock - Current block position
	CurrentBlock() int
	// CurrentTime - Current play time
	// Need to be overridden
	CurrentTime() int
	// resumeLoop - Inner method to start playing loop
	resumeLoop()
}

// tracker - Basic implementation of tracker
type tracker struct {
	Interface
	// trk - Current track
	trk *track.Track
	// pos - Current playing block pos
	pos int
	// playing - Current inner state of playing
	playing bool
}

// resume - Method to start playing loop.
// Override this realisation.
func (t *tracker) resumeLoop() {
	panic("don't use default realisation of resume loop in tracker")
}

// SeekTime - Set cursor position to specific block.
// Override this realisation.
func (t *tracker) SeekTime(pos int) error {
	panic("don't use default realisation of seek time in tracker")
}

// TotalTime - Length in ms of current track
// Override this realisation.
func (t *tracker) TotalTime() int {
	panic("don't use default realisation of total time in tracker")
}

// CurrentTime - Current play time
// Override this realisation.
func (t *tracker) CurrentTime() int {
	panic("don't use default realisation of current time in tracker")
}

// Play - Start async playing
func (t *tracker) Play(trk *track.Track) error {
	panic("don't use default realisation of play in tracker")
}

// Resume - Resume playing
func (t *tracker) Resume() error {
	panic("don't use default realisation of resume in tracker")
}

// Pause - Pause playing
func (t *tracker) Pause() error {
	t.playing = false
	return nil
}

// Stop - Stop playing
func (t *tracker) Stop() error {
	t.playing = false
	t.pos = 0
	return nil
}

// State - Return tracker's state
func (t *tracker) State() State {
	if t.playing {
		return StatePlaying
	}
	if t.pos == 0 {
		return StateStopped
	}
	if t.pos > len(t.trk.Units) {
		return StateFinished
	}
	return StatePaused
}

// SeekBlock - Set cursor position to specific block
func (t *tracker) SeekBlock(pos int) error {
	t.pos = pos
	return nil
}

// TotalBlocks - Count of blocks in current track
func (t *tracker) TotalBlocks() int {
	return len(t.trk.Units)
}

// CurrentBlock - Current block position
func (t *tracker) CurrentBlock() int {
	return t.pos
}