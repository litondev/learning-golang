package main

import (
	"fmt"
	"strings"
)

type student struct{
	name string
	grade int
}

func (s student) sayHello(){
	fmt.Println(s.name)
}

func (s student) getName(i int) string{
	return strings.Split(s.name," ")[i-1]
}

func main(){
	var studentA = student{name:"String"}
	studentA.sayHello()
}