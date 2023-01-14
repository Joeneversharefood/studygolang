package main

import "fmt"

func main() {
	defer func() {
		fmt.Println(0)
	}()
	fmt.Println(1)
}
