package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)
	fmt.Println("starting")
	<-interrupt
	fmt.Println("interrupted")
	fmt.Println("stopping")
}
