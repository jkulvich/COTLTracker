package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"player/cmdline"
	"player/device"
)

func main() {
	cmd, cli := cmdline.Parse()
	_ = cli
	_ = cmd
	fmt.Println(cmd, cli)

	dev, err := device.New()
	if err != nil {
		log.Fatalf("ADB Serv: %s", err)
	}

	fmt.Println(dev.GetVendorModel())
	fmt.Println(dev.GetScreenSizeAlbum())
	fmt.Println(dev.Tap(80, 80))

	// Waiting for shutdown signal
	exitSign := make(chan os.Signal)
	signal.Notify(exitSign, os.Interrupt, os.Kill)
	<-exitSign
}
