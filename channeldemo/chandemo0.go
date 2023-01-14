package main

import (
	"fmt"
	"time"
)

func main() {

	Chan0 := make(chan int)
	go func() {
		fmt.Printf("1\n")
		<-Chan0
		fmt.Printf("2\n")
		fmt.Printf("got signal, I am out\n")
	}()
	fmt.Printf("hahahaha\n")

	for true {
		time.Sleep(20 * time.Second)
		break
	}
	fmt.Printf("dudududdu\n")
	Chan0 <- 1

	time.Sleep(60 * time.Second)
	fmt.Printf("I am main thread, I am out\n")

}
