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
	fmt.Printf("I am goruntine %d,I am waiting for hello,I will sleep 10 sec first\n", myChan.Number)
	time.Sleep(10 * time.Second)
	fmt.Printf("I am goruntine %d,I am wake up\n", myChan.Number)
	x := <-(myChan.Chan)
	if x == "" {
		fmt.Printf("I am goruntine %d,got nothing from chan0, I am Out\n", myChan.Number)
		return
	}
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
	for i := 1; i < 11; i++ {
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
	recv0 := <-myChan0.Chan
	fmt.Printf("main goruntine grab a %s from chan0\n", recv0)

	fmt.Printf("I am main goruntine, I am gonna close chan0\n")
	close(chan0)

	recv1 := <-myChan0.Chan
	fmt.Printf("main goruntine grab a %s from chan0\n", recv1)

	time.Sleep(20 * time.Second)
	fmt.Printf("I am main goruntine, I gonna die\n")

}
