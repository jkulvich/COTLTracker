package main

import (
	"log"
	"os"
	"os/signal"
	"player/cmdline"
	"player/tracker"
	"player/tracker/track"
	"player/tracker/track/parser"
)

func main() {
	cmd, cli := cmdline.Parse()
	_ = cli
	_ = cmd

	if cmd == "play" {
		trk := track.New()

		// Configure specific parser
		var pars parser.Interface
		switch cli.Loader {
		case "cotl":
			pars = parser.NewCOTLTrack()
		}

		// Load specific track
		if err := trk.LoadFile(cli.Play.Track, pars); err != nil {
			log.Fatalf("can't load track: %s", err)
		}

		// Use specific type of tracker
		var t tracker.Interface
		switch cli.Use {
		case "stdout":
			t = tracker.NewStream(os.Stdout, tracker.StreamConfig{
				Tick:  cli.Tick,
				Delay: cli.Delay,
			})
		}

		if err := t.Play(trk); err != nil {
			log.Fatalf("can't start playing: %s", err)
		}
	}

	// Waiting for shutdown signal
	exitSign := make(chan os.Signal)
	signal.Notify(exitSign, os.Interrupt, os.Kill)
	<-exitSign
}
