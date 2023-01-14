package main

import (
    "fmt"
    "time"
)

func eat(s string){
    for true{
        time.Sleep(1 * time.Second)
        fmt.Printf("eating %s\n", s) 
    }
}


func main(){
    
    apple := [...]byte{'a','p','p','l','e'}
    banana := [...]byte{'b','a','n','a','n','a'}
    
    arg := string(apple[:])

    go eat(string(arg))
    for true {
        time.Sleep(1 * time.Second)
        fmt.Printf("eating %s\n", banana)
    }
}
