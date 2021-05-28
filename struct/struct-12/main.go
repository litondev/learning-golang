package main

import (
	"struct-12/library"
	"fmt"
)

func main(){
	library.SayHello()
	
	var studentA = library.Student{Name:"ss"}
	fmt.Println(studentA.Name)
	// fmt.Println(studentA.grade)
}