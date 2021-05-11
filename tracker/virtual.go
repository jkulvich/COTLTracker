package tracker

import (
	"bufio"
	"fmt"
	"github.com/dbatbold/beep"
	"io"
	"io/ioutil"
	"player/tracker/track"
	"player/tracker/track/unit"
	"strings"
	"time"
)

// Virtual - Virtual tracker.
// It plays song on your desktop device.
type Virtual struct {
	Interface
	// conf - Tracker configuration
	conf VirtualConfig
	// trk - Current track
	trk *track.Track
	// pos - Current playing block pos
	pos int
	// playing - Current inner state of playing
	playing bool
}

// VirtualConfig - Virtual tracker configuration
type VirtualConfig struct {
	// Log - If presented - print additional info
	Log *io.Writer
	// Tick - Length of tick in ms. Default 100
	Tick int
	// Delay - Length of delay between taps. Default 0
	Delay int
}

// NewVirtual - Create new virtual tracker
func NewVirtual(config VirtualConfig) *Virtual {
	if config.Log == nil {
		config.Log = &ioutil.Discard
	}

	tracker := Virtual{
		conf: config,
	}

	tracker.log("!!! Virtual tracker is in dev mode !!!")

	return &tracker
}

// Play - Start async playing
func (t *Virtual) Play(trk *track.Track) error {
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
func (t *Virtual) Pause() error {
	t.playing = false
	return nil
}

// Resume - Resume playing
func (t *Virtual) Resume() error {
	t.playing = true
	go func() {
		t.resumeLoop()
		t.playing = false
	}()
	return nil
}

// Stop - Stop playing
func (t *Virtual) Stop() error {
	t.playing = false
	t.pos = 0
	return nil
}

// State - Return tracker's state
func (t *Virtual) State() State {
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
func (t *Virtual) SeekBlock(pos int) error {
	t.pos = pos
	return nil
}

// SeekTime - Set cursor position to specific block at time.
// Override this realisation.
func (t *Virtual) SeekTime(pos int) error {
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
func (t *Virtual) TotalBlocks() int {
	if t.trk != nil {
		return len(t.trk.Units)
	}
	return 0
}

// TotalTime - Length in ms of current track
// Override this realisation.
func (t *Virtual) TotalTime() int {
	return t.timeOf(t.trk.Len())
}

// CurrentTime - Current play time
// Override this realisation.
func (t *Virtual) CurrentTime() int {
	return t.timeOf(t.pos)
}

// CurrentBlock - Current block position
func (t *Virtual) CurrentBlock() int {
	return t.pos
}

// resumeLoop - Start loop of playing
func (t *Virtual) resumeLoop() {
	// Preparing device and music object
	_ = beep.OpenSoundDevice("default")
	defer beep.CloseSoundDevice()
	music := beep.NewMusic("")

	for {
		if !t.playing || t.pos >= len(t.trk.Units) {
			t.playing = false
			break
		}
		// Create an log stub
		logHeader := fmt.Sprintf("%5.1f%% [%4d/ %4d blk] [%6d/ %6d ms]",
			float64(t.pos)/float64(t.TotalBlocks())*100,
			t.pos,
			t.TotalBlocks(),
			t.CurrentTime(),
			t.TotalTime(),
		)
		// Fetch current unit
		u := t.trk.Units[t.pos]
		// Play note
		if u.Type == unit.TypeNote {
			t.log("%s %s", logHeader, u.Note.String())
			go t.playNote(music, u.Note)
			time.Sleep(time.Millisecond * time.Duration(t.conf.Delay))
		}
		// Delay
		if u.Type == unit.TypeDelay {
			t.log("%s %s", logHeader, strings.Repeat("-", int(u.Delay)))
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
func (t *Virtual) timeOf(pos int) int {
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

func (t *Virtual) playNote(music *beep.Music, note unit.Note) {
	// Encoder from notes to beep format
	noteMap := map[string]string{
		"C0": "q", "D0": "w", "E0": "e", "F0": "r", "G0": "t",
		"A0": "y", "B0": "u", "C1": "i", "D1": "o", "E1": "p",
		"F1": "[", "G1": "]", "A1": "z", "B1": "x", "C2": "c",
	}

	// Beep string
	beepCode := fmt.Sprintf("VD DE %s", noteMap[note.String()])

	// Playing
	_ = beep.InitSoundDevice()
	go music.Play(bufio.NewReader(strings.NewReader(beepCode)), 100)
	music.Wait()
	beep.FlushSoundBuffer()
}

// log - Inner verbose logger
func (t *Virtual) log(str string, a ...interface{}) {
	_, _ = io.WriteString(*t.conf.Log, fmt.Sprintf(str+"\n", a...))
}
