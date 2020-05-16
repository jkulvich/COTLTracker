package main

import (
	"flag"
	"io/ioutil"
	"log"
	"player/cotl"
	"time"
)

func main() {

	flagTrack := flag.String("track", "", "path to track file")
	flagDelay := flag.Int("delay", 80, "min delay between taps")
	flagStart := flag.Int("start", 0, "start block position")
	flagTest := flag.Bool("test", false, "make a sound test")
	flag.Parse()

	var stave []byte

	if *flagTest {
		stave = []byte(`C1 - D1 - E1 - F1 - G1 - A1 - B1 - C2 - D2 - E2 - F2 - G2 - A2 - B2 - C3`)
	} else {
		var err error
		stave, err = ioutil.ReadFile(*flagTrack)
		if err != nil {
			log.Fatal(err)
		}
	}

	// Парсинг нового трека
	track, err := cotl.NewTrack(string(stave))
	if err != nil {
		log.Fatal(err)
	}

	// Создание нового трекер и подключение к устройству
	tracker, err := cotl.New(*flagDelay)
	if err != nil {
		log.Fatal(err)
	}

	// Старт воспроизведения композиции
	if err := tracker.Play(track, *flagStart); err != nil {
		log.Fatal(err)
	}

	<-time.After(time.Millisecond * 1000)
}
