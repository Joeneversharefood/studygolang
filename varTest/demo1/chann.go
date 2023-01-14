package main

import "fmt"

var ch chan string

func main() {

	fmt.Println(ch)
	fmt.Println(&ch)
	ch = make(chan string)
	fmt.Println(ch)
	fmt.Println(&ch)

}
