package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

var onceLoader sync.Once

var num int = 0

func add() {
	defer wg.Done()
	num++
}

func reporter() {
	defer wg.Done()
	for {
		fmt.Println(num)
		time.Sleep(time.Second)
	}
}

func main() {

	for i := 0; i < 10; i++ {
		wg.Add(1)
		fmt.Println("add one adder")
		go func() { onceLoader.Do(add) }()
	}
	wg.Add(1)
	go reporter()
	wg.Wait()
}
