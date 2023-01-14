package main

import (
    "fmt"
    "net/http"
)

func routepagehandler(w http.ResponseWriter, r *http.Request){
    if r.URL.RequestURI() == "/favicon.ico"{
        return
    }
    fmt.Fprintf(w, "Hello~\n")
    fmt.Printf("hahahaha\n")
}

func main(){
    http.HandleFunc("/",routepagehandler)
    http.ListenAndServe("localhost:8000",nil)
}

