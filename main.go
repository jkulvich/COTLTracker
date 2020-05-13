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
	flagDelay := flag.Int("delay", 50, "min delay between taps")
	flagStart := flag.Int("start", 0, "start block position")
	flag.Parse()

	stave, err := ioutil.ReadFile(*flagTrack)
	if err != nil {
		log.Fatal(err)
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
