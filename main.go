package main

import (
	"flag"
	"io/ioutil"
	"log"
	"player/cotltracker"
	"time"
)

func main() {

	flagSerial := flag.String("serial", "", "ADB smartphone serial id")
	flagTrack := flag.String("track", "", "path to track file")
	flagSpeed := flag.Float64("speed", 1, "track playing speed")
	flag.Parse()

	stave, err := ioutil.ReadFile(*flagTrack)
	if err != nil {
		log.Fatal(err)
	}

	track, err := cotltracker.NewTrack(string(stave))
	if err != nil {
		log.Fatal(err)
	}

	tracker, err := cotltracker.New(*flagSerial)
	if err != nil {
		log.Fatal(err)
	}

	if err := tracker.Play(track, float32(*flagSpeed)); err != nil {
		log.Fatal(err)
	}

	<-time.After(time.Millisecond * 1000)
}
