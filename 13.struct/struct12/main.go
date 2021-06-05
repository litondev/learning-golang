package main

import (
	"struct12/library"
	"fmt"
)

func main(){
	library.SayHello()
	
	var studentA = library.Student{Name:"Nama"}
	fmt.Println(studentA.Name)
	fmt.Println(studentA.grade)
}