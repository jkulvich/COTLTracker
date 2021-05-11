package main

import (
	"github.com/faiface/beep"
	"github.com/faiface/beep/speaker"
	"log"
	"math"
	"os"
	"os/signal"
	"player/cmdline"
	"player/tracker"
	"player/tracker/track"
	"player/tracker/track/parser"
	"time"
)

func Noise() beep.Streamer {
	return beep.StreamerFunc(func(samples [][2]float64) (n int, ok bool) {
		for i := range samples {
			p := math.Sin(float64(i * 4))
			samples[i][0] = p
			samples[i][1] = p
		}
		return len(samples), true
	})
}

func main() {
	// CLI
	cmd, cli := cmdline.Parse()

	sr := beep.SampleRate(44100)
	speaker.Init(sr, sr.N(time.Second/10))
	speaker.Play(beep.Seq(beep.Take(sr.N(1*time.Second), Noise())))

	// Exit signal
	exitSign := make(chan os.Signal)

	if cmd == "play" {
		trk := track.New()

		// Configure specific parser
		var pars parser.Interface
		switch cli.Play.Loader {
		case "cotl":
			pars = parser.NewCOTLTrack()
		}

		// Load specific track
		if err := trk.LoadFile(cli.Play.Track, pars); err != nil {
			log.Fatalf("can't load track: %s", err)
		}

		// Make new slice from original track
		trk = trk.Sub(cli.Play.Start, trk.Len())

		// Find timing value in track comments or CLI
		timing := trk.GetTiming()
		if cli.Play.Tick != -1 {
			timing = cli.Play.Tick
		}
		if timing == -1 {
			timing = 200
		}

		// Normalize track and shift by shift comment
		trk.Normalize()
		trk.Shift(trk.GetShift())

		// Mod specific type of tracker
		var t tracker.Interface
		switch cli.Play.Mod {
		case "stdout":
			t = tracker.NewStream(os.Stdout, tracker.StreamConfig{
				Tick:  timing,
				Delay: cli.Play.Delay,
			})
		case "report":
			t = tracker.NewReport(tracker.ReportConfig{
				Tick:  timing,
				Delay: cli.Play.Delay,
			})
		}

		// Start play
		if err := t.Play(trk); err != nil {
			log.Fatalf("can't start playing: %s", err)
		}

		// Wait for finish
		go func() {
			for {
				if t.State() == tracker.StateFinished {
					exitSign <- os.Interrupt
				}
				time.Sleep(time.Millisecond * 100)
			}
		}()
	}

	// Waiting for shutdown signal
	signal.Notify(exitSign, os.Interrupt, os.Kill)
	<-exitSign
}
