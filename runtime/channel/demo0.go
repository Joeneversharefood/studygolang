package main

import (
	"fmt"
	"sync"
)

var intChan chan int

var wg sync.WaitGroup

func addToChan(i int) {
	defer wg.Done()
	fmt.Printf("send %d into chan\n", i)
	intChan <- i
}

func grabItemFromChan(i int) {
	defer wg.Done()
	num := <-intChan
	fmt.Printf("rountine[%d] grab %d from channel\n", i, num)
}

func main() {
	intChan = make(chan int)
	wg.Add(3)
	go addToChan(0)
	go addToChan(1)
	go addToChan(2)

	for val := range intChan {
		fmt.Println(val)
	}

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go grabItemFromChan(i)
	}
	wg.Wait()

}
