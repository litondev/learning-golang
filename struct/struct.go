package main 

import "fmt"

func main(){
	type student struct {
		name string
		grade int
	}

	var classA = student{}
	classA.name = "hai"
	classA.grade = 10

	fmt.Println(classA.name)
	fmt.Println(classA.grade)

	var classB = student{name: "hai",grade: 3}
	fmt.Println(classB.name)
	fmt.Println(classB.grade)
}