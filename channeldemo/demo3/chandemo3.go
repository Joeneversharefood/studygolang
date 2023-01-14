package main

import (
	"fmt"
	"time"
)

var chanqueuenum int = 10

type mychan struct {
	Number int
	Chan   chan string
}

func (myChan *mychan) recv() {
	fmt.Printf("I am goruntine %d,I am waiting for hello\n", myChan.Number)
	x := <-(myChan.Chan)
	fmt.Printf("I am goruntine %d,recv = %s,Out\n", myChan.Number, x)
}

type sendmessagetochan interface {
	send(string)
}

func (myChan *mychan) send(str string) {

	fmt.Printf("send %s to chan\n", str)
	myChan.Chan <- str
}

func main() {

	chan0 := make(chan string, chanqueuenum)
	for i := 1; i < 5; i++ {
		myChan := &mychan{Chan: chan0, Number: i}
		go myChan.recv()
	}

	time.Sleep(5 * time.Second)

	myChan0 := &mychan{Chan: chan0, Number: 0}

	var str string
	for j := 0; j < 10; j++ {
		str = fmt.Sprintf("hello%d", j)
		myChan0.Chan <- str
		fmt.Printf("send to chan : %s,&str = %v\n", str, &str)
	}

	time.Sleep(5 * time.Second)

}
