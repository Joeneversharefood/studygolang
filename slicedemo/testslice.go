package main

import "fmt"

func main(){
   

    var arr0 = [...]int{1,2,3,4,5,6} 
    var slice0 = arr0[2:4]

    var slice1 []byte = make([]byte, 32, 64)

    slice1[0] = '1'
    
    var slice2 []string = []string{"yuan", "ye", "shi", "S", "B"}
    
    fmt.Printf("slice0 len = %d\tcap = %d\n",len(slice0), cap(slice0))
    fmt.Println(slice0)

    for i, val := range(slice0){
        fmt.Printf("slice0 index = %d\taddr = %v\n", i, val)
    }
    
    slice0 = append(slice0, 1, 1)

    fmt.Println("slice0 after apend one")

    for i, val := range(slice0){
        fmt.Printf("slice0 index = %d\taddr = %v\n", i, val)                                                                                                                                                                                                                                  
    }

    fmt.Printf("slice0 len = %d\tcap = %d\n",len(slice0), cap(slice0))
    fmt.Println(slice0)

    fmt.Println("arr:\n")
    fmt.Println(arr0)

    fmt.Printf("slice1 len = %d\tcap = %d\n",len(slice1), cap(slice1))
    fmt.Println(slice1)
    fmt.Printf("slice2 len = %d\tcap = %d\n",len(slice2), cap(slice2))
    fmt.Println(slice2)
}

