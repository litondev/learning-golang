package main

import "fmt"

type student struct {
	class struct{
		name string
	}

	grade int
}

func main(){
	var studentA = student{}
	studentA.class.name = "a"
	fmt.Println(studentA)
}