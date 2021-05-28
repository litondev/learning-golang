package main

import "fmt"


func main(){
	var student struct {
		name string
		age int
	}

	student.name = "Testing"
	fmt.Println(student.name)

	var mystudent struct {
		name string 
		age int
	}

	mystudent.name = "test"

	var stillstudent = struct {
		grade int
	}{39}

	fmt.Println(mystudent.name)
	fmt.Println(stillstudent.grade)
}