package main

import (
	"fmt"
	"strconv"
)



func main(){
	person:= new(Person)
	person.Name = "Tinuade"
	person.Age = 22

fmt.Println(person.getPersonInfo())
}

type Person struct{
	Name string
	Age uint
}

func(p Person) getPersonInfo()string{
	return p.Name + " "+ strconv.Itoa(int(p.Age))
}