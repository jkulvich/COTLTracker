package main

import (
	"flag"
	"io/ioutil"
	"log"
	"player/cotl"
	"time"
)

func main() {

	flagSerial := flag.String("serial", "", "ADB smartphone serial id")
	flagTrack := flag.String("track", "", "path to track file")
	flagSpeed := flag.Float64("speed", 1, "track playing speed")
	flagAdb := flag.String("adb", "adb", "path where ADB tool located")
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
	tracker, err := cotl.New(*flagAdb, *flagSerial)
	if err != nil {
		log.Fatal(err)
	}

	// Старт воспроизведения композиции
	if err := tracker.Play(track, float32(*flagSpeed)); err != nil {
		log.Fatal(err)
	}

	<-time.After(time.Millisecond * 1000)
}
