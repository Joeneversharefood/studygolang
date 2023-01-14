package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func sayHello(i int) {
	fmt.Printf("enter routine[%d]\n", i)
	defer wg.Done()
	defer fmt.Printf("leave rountine[%d]0\n", i)
	defer fmt.Printf("leave rountine[%d]1\n", i)
	defer fmt.Printf("leave rountine[%d]2\n", i)
	runtime.Gosched() //注释与不注释就能看出来释放cpu的效果
	fmt.Printf("routine[%d]：0\n", i)
	fmt.Printf("routine[%d]：0\n", i)
}

func main() {
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go sayHello(i)
	}
	wg.Wait()
}
