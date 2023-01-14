package main

import (
	"fmt"
	"reflect"
)

type demo0 struct {
}

func (d *demo0) f0(a, b int) {
	fmt.Println(a, b)
}

func main() {
	test0 := demo0{}
	type0 := reflect.TypeOf(&test0)
	value0 := reflect.TypeOf(&test0)

	fmt.Printf("type : %v,value : %v\n", type0, value0)

	fmt.Printf("method num = %d\n", type0.NumMethod())
	for i := 0; i < type0.NumMethod(); i++ {
		fmt.Printf("method[%d]= %v args : %d\n", i, type0.Method(i).Name, type0.Method(i).Type.NumIn())
	}

}
