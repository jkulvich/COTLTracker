package tracker

import (
	"bufio"
	"fmt"
	"github.com/dbatbold/beep"
	"player/tracker/track/unit"
	"strings"
	"time"
)

// Virtual - Virtual tracker.
// It plays song on your desktop device.
type Virtual struct {
	Simple
	// music - Current music object
	music *beep.Music
}

// VirtualConfig - Virtual tracker configuration
type VirtualConfig struct {
	// Tick - Length of tick in ms. Default 100
	Tick int
	// Delay - Length of delay between taps. Default 0
	Delay int
}

// NewVirtual - Create new virtual tracker
func NewVirtual(config VirtualConfig) *Simple {
	tracker := &Virtual{
		music: beep.NewMusic(""),
	}
	tracker.Delay = config.Delay
	tracker.Tick = config.Tick
	return Simplify(tracker)
}

func (t *Virtual) beforePlay(data *Simple) error {
	return beep.OpenSoundDevice("default")
}

func (t *Virtual) afterPlay(data *Simple) error {
	beep.CloseSoundDevice()
	return nil
}

func (t *Virtual) step(data *Simple, u unit.Unit) error {
	// Create an log stub
	logHeader := fmt.Sprintf("%5.1f%% [%4d/ %4d blk] [%6d/ %6d ms]",
		float64(data.pos)/float64(data.TotalBlocks())*100,
		data.pos,
		data.TotalBlocks(),
		data.CurrentTime(),
		data.TotalTime(),
	)
	// Play note
	if u.Type == unit.TypeNote {
		fmt.Printf("%s %s\n", logHeader, u.Note.String())
		go t.playNote(t.music, u.Note)
		time.Sleep(time.Millisecond * time.Duration(t.Delay))
	}
	// Delay
	if u.Type == unit.TypeDelay {
		fmt.Printf("%s %s\n", logHeader, strings.Repeat("-", int(u.Delay)))
		time.Sleep(time.Millisecond * time.Duration(t.Tick*int(u.Delay)))
	}
	return nil
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
