package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

var rwLock sync.RWMutex

var jobs int = 3

func doJob(i int) {
	defer wg.Done()
	defer rwLock.Unlock()

	rwLock.Lock()
	if jobs > 0 {
		jobs--
		fmt.Printf("rountine %d finish one job\n", i)
	}

}

func reporter() {
	defer wg.Done()

	for {
		rwLock.RLock()
		fmt.Printf("there is %d job left\n", jobs)
		rwLock.RUnlock()
		time.Sleep(500 * time.Millisecond)
	}
}

func main() {

	wg.Add(1)

	go reporter()

	time.Sleep(time.Second)

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go doJob(i)
		time.Sleep(200 * time.Millisecond)
	}

	wg.Wait()
}
