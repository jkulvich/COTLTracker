package tracker

import (
	"player/tracker/track"
	"player/tracker/track/unit"
	"time"
)

// ManagedInterface - Another tracker's interface with reduced set of methods.
type ManagedInterface interface {
	// step - Play unit step
	step(n unit.Unit) error
}

// ManagedConfig - Basic realisation of config
type ManagedConfig struct {
	// Tick - Length of tick in ms for delay nodes like 'dashes'
	Tick int
	// Delay - Length of delay between taps to prevent taps throttling
	Delay int
}

// Managed - Basic realisation of managed wrapper
type Managed struct {
	Interface
	// tracker - Actual managed object
	tracker ManagedInterface
	// config - Base configuration
	config ManagedConfig
	// trk - Current track
	trk track.Track
	// pos - Current playing block pos
	pos int
	// playing - Current inner state of playing
	playing bool
}

// Manage - Takes some work of managing trackers.
// It can be helpful for describing simple trackers which doesn't have complicated functionality.
// So, you can use this to simplify your tracker realisation.
func Manage(tracker ManagedInterface) *Managed {
	return &Managed{
		tracker: tracker,
	}
}

// Play - Start async playing
func (t *Managed) Play(trk track.Track) error {
	t.trk = trk
	t.pos = 0
	if err := t.Resume(); err != nil {
		return err
	}
	return nil
}

// Pause - Pause playing
func (t *Managed) Pause() error {
	t.playing = false
	return nil
}

// Resume - Resume playing
func (t *Managed) Resume() error {
	t.playing = true
	go func() {
		t.resumeLoop()
		t.playing = false
	}()
	return nil
}

// Stop - Stop playing
func (t *Managed) Stop() error {
	t.playing = false
	t.pos = 0
	return nil
}

// State - Return tracker's state
func (t *Managed) State() State {
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
func (t *Managed) SeekBlock(pos int) error {
	t.pos = pos
	return nil
}

// SeekTime - Set cursor position to specific block at time.
// Override this realisation.
func (t *Managed) SeekTime(pos int) error {
	var total time.Duration
	for i, u := range t.trk.Units {
		switch u.Type {
		case unit.TypeDelay:
			total += time.Millisecond * time.Duration(u.Delay) * time.Duration(t.config.Tick)
		case unit.TypeNote:
			total += time.Millisecond * time.Duration(t.config.Delay)
		}
		if pos < int(total.Milliseconds()) {
			t.pos = i
			return nil
		}
	}
	return nil
}

// TotalBlocks - Count of blocks in current track
func (t *Managed) TotalBlocks() int {
	return len(t.trk.Units)
}

// TotalTime - Length in ms of current track
// Override this realisation.
func (t *Managed) TotalTime() int {
	return t.timeOf(t.trk.Len())
}

// CurrentTime - Current play time
// Override this realisation.
func (t *Managed) CurrentTime() int {
	return t.timeOf(t.pos)
}

// CurrentBlock - Current block position
func (t *Managed) CurrentBlock() int {
	return t.pos
}

// resumeLoop - Start loop of playing
func (t *Managed) resumeLoop() error {
	for {
		if !t.playing || t.pos >= len(t.trk.Units) {
			t.playing = false
			break
		}
		// Play note
		if err := t.tracker.step(t.trk.Units[t.pos]); err != nil {
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
func (t *Managed) timeOf(pos int) int {
	var total time.Duration
	for i := 0; i < pos; i++ {
		u := t.trk.Units[i]
		switch u.Type {
		case unit.TypeDelay:
			total += time.Millisecond * time.Duration(u.Delay) * time.Duration(t.config.Tick)
		case unit.TypeNote:
			total += time.Millisecond * time.Duration(t.config.Delay)
		}
	}
	return int(total.Milliseconds())
}
