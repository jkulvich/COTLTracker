package main

import (
	"io/ioutil"
	"log"
	"os"
	"os/signal"
	"player/cmdline"
	"player/tracker"
	"player/tracker/track"
	"player/tracker/track/parser"
	"strings"
	"time"
)

func main() {
	// CLI
	cmd, cli := cmdline.Parse()

	// Exit signal
	exitSign := make(chan os.Signal)

	if cmd == "play <track>" {
		trk := track.New()

		// Configure specific parser
		var pars parser.Interface
		switch cli.Play.Loader {
		case "cotl":
			pars = parser.NewCOTLTrack()
		}

		// Load specific track
		if err := loadTrack(trk, cli.Play.Track, pars); err != nil {
			log.Fatalf("can't load track: %s", err)
		}

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

		// Additional logger
		verboseLog := ioutil.Discard
		if cli.Play.Verbose == true {
			verboseLog = os.Stdout
		}

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
		case "virtual":
			t = tracker.NewVirtual(tracker.VirtualConfig{
				Tick:  timing,
				Delay: cli.Play.Delay,
				Log: &verboseLog,
			})
		default:
			panic("can't find module: " + cli.Play.Mod)
		}

		// Start play
		if err := t.Play(trk); err != nil {
			log.Fatalf("can't start playing: %s", err)
		}

		// Set specific position
		_ = t.SeekBlock(cli.Play.Start)

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

func loadTrack(trk *track.Track, filename string, pars parser.Interface) error {
	// Add default filename
	if !strings.HasSuffix(filename, ".txt") {
		filename += ".txt"
	}
	// If URL try to load
	if strings.HasPrefix(filename, "http://") || strings.HasPrefix(filename, "https://") {
		if err := trk.LoadURL(filename, pars); err != nil {
			return err
		}
	} else {
		// In other case it is a file path
		if err := trk.LoadFile(filename, pars); err != nil {
			baseURL := "https://raw.githubusercontent.com/jkulvich/COTLTracker/master/tracks/"
			if err := trk.LoadURL(baseURL + filename, pars); err != nil {
				return err
			}
		}
	}
	return nil
}