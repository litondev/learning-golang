package main

import "fmt"

type student struct{
	name string
	grade int
}

func (s *student) sayHello(){
	fmt.Println(s.name)
}

func main(){
	var studentA = student{name:"Sd"}
	studentA.sayHello()	
}