package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"player/cmdline"
	"player/services/adb"
)

func main() {
	cmd, cli := cmdline.Parse()
	_ = cli
	_ = cmd
	fmt.Println(cmd, cli)

	serviceADB, err := adb.NewADBService()
	if err != nil {
		log.Fatalf("ADB Serv: %s", err)
	}

	fmt.Println(serviceADB.GetVendorModel())
	fmt.Println(serviceADB.GetScreenSizeAlbum())
	fmt.Println(serviceADB.Tap(80, 80))

	// Waiting for shutdown signal
	exitSign := make(chan os.Signal)
	signal.Notify(exitSign, os.Interrupt, os.Kill)
	<-exitSign
}
