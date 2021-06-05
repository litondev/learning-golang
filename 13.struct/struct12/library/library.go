package library

import "fmt"

// cannot be access from outside
// type student struct{
type Student struct{
	Name string
	grade int
}

func SayHello(){
	fmt.Println("Testing")
}