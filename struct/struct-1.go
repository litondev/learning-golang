package main

import "fmt"

func main() {
	type student struct {
		name string
		grade int
	}

	var classA = student{}
	classA.name = "hai"
	classA.grade = 10

	var classB *student = &classA

	fmt.Println(classB.name)

}