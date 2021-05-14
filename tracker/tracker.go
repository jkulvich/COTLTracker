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
// Play is an async, so you can continue just after call.
type Interface interface {
	// Play - Start async play
	Play(trk track.Track)
	// Pause - Pause playing
	Pause()
	// Resume - Resume async play
	Resume()
	// Stop - Pause and set cursor position to start
	Stop()
	// State - Tracker's current state like StateStopped, StatePlaying, StatePaused and StateFinished
	State() State
	// SeekBlock - Set cursor position to specific block
	SeekBlock(pos int)
	// SeekTime - Set cursor position to specific block at track time.
	SeekTime(ms int)
	// TotalBlocks - Count of blocks in current track
	TotalBlocks() int
	// TotalTime - Length in ms of current track
	TotalTime() int
	// CurrentBlock - Current block position
	CurrentBlock() int
	// CurrentTime - Current play time
	CurrentTime() int
	// Error - Return last error happens in the tracker
	Error(flush ...bool) error
}
