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
// Tracker using to Manage track playing.
// Play and Resume should be async functions and shouldn't block current goroutine
type Interface interface {
	// Play - Start async play
	// Need to be overridden
	Play(trk track.Track) error
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
}