package main

import (
	"log"
	"os"
	"os/signal"
	"player/cmdline"
	"player/tracker"
	"player/tracker/track"
	"player/tracker/track/parser"
	"time"
)

func main() {
	// CLI
	cmd, cli := cmdline.Parse()

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
		case "info":
			t = tracker.NewInfo(os.Stdout, tracker.InfoConfig{
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
