package main

import "fmt"

func main(){
    

    var slice0 []int = []int{0,1,2,3,4,5,6,7}
    
    var arr0 = [...]int{1,1,1,1,1,1,1,1,1,1}

    slice1 := arr0[:4]

    
    

    fmt.Printf("slice0[0] = %v\t slice0[1]=%v\n",&slice0[0], &slice0[1])

    fmt.Printf("slice1[0] = %v\t slice1[1]=%v\n",&slice1[0], &slice1[1])

    ptr0 := &slice1[0]
    ptr1 := &slice1[1]
    
    fmt.Printf("ptr0 = %v\tptr1 = %v\n", ptr0, ptr1)

    fmt.Printf("*ptr0 = %v\t*ptr1 = %v\n", *ptr0, *ptr1)

    *ptr0 = 2
    *ptr1 = 3

    fmt.Printf("ptr0 = %v\tptr1 = %v\n", ptr0, ptr1)
    fmt.Printf("*ptr0 = %v\t*ptr1 = %v\n", *ptr0, *ptr1)
    
    
    fmt.Printf("slice1 len = %d\t cap = %d\n", len(slice1), cap(slice1))
    
    slice1 = append(slice1, 9, 9, 9)

    fmt.Printf("append(slice1, 9, 9, 9)\n")

    
    


    fmt.Println(slice1)

    fmt.Printf("slice1[0] = %v\t slice1[1]=%v\n",&slice1[0], &slice1[1])

    fmt.Printf("slice1 len = %d\t cap = %d\n", len(slice1), cap(slice1))

    slice1 = append(slice1, 9, 9, 9)

    fmt.Printf("append(slice1, 9, 9, 9)\n")
    fmt.Println(slice1)
    fmt.Printf("slice1[0] = %v\t slice1[1]=%v\n",&slice1[0], &slice1[1])
 
    fmt.Printf("slice1 len = %d\t cap = %d\n", len(slice1), cap(slice1))
    
    slice1 = append(slice1, 9)

    fmt.Printf("append(slice1, 9)\n")
    fmt.Println(slice1)
    fmt.Printf("slice1[0] = %v\t slice1[1]=%v\n",&slice1[0], &slice1[1])
    fmt.Printf("slice1 len = %d\t cap = %d\n", len(slice1), cap(slice1))
    
    fmt.Printf("ptr0 = %v\tptr1 = %v\n", ptr0, ptr1)
    

    fmt.Println()
    fmt.Println()
    fmt.Printf("slice0[0] = %v\t, slice0[1] = %v\n", &slice0[0], &slice0[1])
    fmt.Printf("slice0 len = %d\t cap = %d\n", len(slice0), cap(slice0))
    slice0 = append(slice0, 1, 1)
    
    fmt.Printf("slice0[0] = %v\t, slice0[1] = %v\n", &slice0[0], &slice0[1])
    fmt.Printf("slice0 len = %d\t cap = %d\n", len(slice0), cap(slice0))
    
}

