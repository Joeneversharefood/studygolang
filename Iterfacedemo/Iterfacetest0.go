package main

import "fmt"

type Animal interface{
    eat()
    sleep()
}

type Dog struct{
    name string
}

type Human struct{
    height float32
    weight float32
}

type Cat struct{
    sleephour uint32
}

func (dog Dog)eat(){
    fmt.Printf("I am a dog named\t%s,I like eat bone\n",dog.name)
}

func (dog Dog)sleep(){
    fmt.Printf("our dog usually sleep 10 hours in a single day\n")
}

func (cat Cat)eat(){
    fmt.Printf("I am a cat,I like eat fish\n")
}

func (cat Cat)sleep(){
    fmt.Printf("I am a cat, I usually sleep\t%d in a single day\n",cat.sleephour)
}

func (human Human)eat(){
    fmt.Printf("I am a human being, I like eat fruit\n")
}

func (human Human)sleep(){
    fmt.Printf("I am a human being. I usually sleep 8 hours in a single day\n")
}


func introduceSelf(a Animal){
    a.eat()
    a.sleep()
}

func main(){
    
    cat0 := new(Cat)
    cat0.sleephour = 12

    dog0 := new(Dog)
    dog0.name = "pat"

    human0 := new(Human)
    human0.height = 172.5
    human0.weight = 132.5
    
    introduceSelf(cat0)
    fmt.Println() 
    introduceSelf(dog0)
    fmt.Println() 
    introduceSelf(human0)
    

}

