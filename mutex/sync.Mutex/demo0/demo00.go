package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

var lock sync.Mutex

var jobs int = 3

func doJob(i int) {
	defer wg.Done()
	defer lock.Unlock()

	lock.Lock()
	if jobs > 0 {
		jobs--
		fmt.Printf("rountine %d finish one job,now have %d left\n", i, jobs)
	} else {
		fmt.Printf("rountine %d there is no more jobs to do\n", i)
	}

}

func main() {

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go doJob(i)
	}

	wg.Wait()
}
