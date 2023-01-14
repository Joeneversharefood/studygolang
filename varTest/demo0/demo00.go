package main

import "fmt"

var a int

var b string

var c map[string]string

func main() {

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	if b == "" {
		fmt.Println("hahahha")
	}
	fmt.Println(1111)

	fmt.Println(&a)
	fmt.Println(&b)
	fmt.Println(&b)

}
