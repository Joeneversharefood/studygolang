package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"
)

func main() {

	sigChan := make(chan os.Signal)
	go func() {
		fmt.Printf("1\n")
		<-sigChan
		fmt.Printf("2\n")
		fmt.Printf("got signal, I am out\n")
	}()
	fmt.Printf("hahahaha\n")

	for true {
		time.Sleep(20 * time.Second)
		break
	}
	fmt.Printf("dudududdu\n")
	signal.Notify(sigChan, os.Interrupt)

	time.Sleep(60 * time.Second)
	fmt.Printf("I am main thread, I am out\n")

}
