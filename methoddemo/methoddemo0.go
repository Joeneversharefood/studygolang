package main

import "fmt"


type Animal struct{
    food,color,name string
    age,weight  uint32
}



type Duck struct{
    
    Animal
    eggnum uint32
}

type Frog struct{
    
    Animal
    height uint32
}

type Behavior interface{
    call()
    eat()
}

func (duck Duck)call(){
    fmt.Printf("I am a duck, named\t%s,my voice is gagagagagagaga~\n", duck.name)
}

func (duck Duck)eat(){
    fmt.Printf("I am a duck, I like eat\t%s\n",duck.food)
}

func (frog Frog)call(){
    fmt.Printf("I am a frog, named\t%s,my voice is guaguaguaguagua~\n", frog.name)
}

func (frog Frog)eat(){
    food := frog.food
    fmt.Printf("I am a frog, I like eat\t%s\n",food)
}

func main(){
    var behavior Behavior
    var behavior0 Behavior
    frog0 := Frog{Animal : Animal{food : "wheat", color : "yellow", name : "jack", age : 1, weight : 12}, height : 30}
    duck0 := Duck{Animal : Animal{food : "wheat", color : "yellow", name : "bob",  weight : 12}, eggnum : 3}
    behavior = frog0
    behavior0 = duck0
    
    behavior.call()
    fmt.Printf("behavior type: %V\n",behavior)
    fmt.Println()
    behavior.eat()
    fmt.Println()
    behavior0.call()
    fmt.Printf("behavior0 type: %V\n",behavior0)
    fmt.Println()
    behavior0.eat()
    frog0.call()
    fmt.Printf("frog0 type = %V\t duck0 type = %V\n",frog0,duck0)
}

