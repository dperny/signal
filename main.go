package main

import (
	"fmt"
	"os"
	"os/signal"
)

func main() {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)
	fmt.Println("starting")
	<-interrupt
	fmt.Println("interrupted")
	fmt.Println("stopping")
}
