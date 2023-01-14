package main

import (
    "fmt"
    "os"
)
func main(){
    
    args := os.Args
    argnum := len(os.Args)

    for i := 0; i < argnum ; i++{
        fmt.Printf("args[%d]\t=\t%s\n",i,args[i])
        fmt.Println()
    }

}

