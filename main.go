package main

import (
	"fmt"
	"os"
	"os/signal"
	"player/cmdline"
)

func main() {
	cmd, cli := cmdline.Parse()
	_ = cli
	_ = cmd
	fmt.Println(cmd, cli)

	//fmt.Println(color.Colorize(color.Red, "It's OK"))

	// Waiting for shutdown signal
	exitSign := make(chan os.Signal)
	signal.Notify(exitSign, os.Interrupt, os.Kill)
	<-exitSign
}
