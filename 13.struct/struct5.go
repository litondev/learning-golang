package main

import "fmt"


func main(){
	var student struct {
		name string
		age int
	}
	student.name = "Nama"
	fmt.Println(student.name)

	var mystudent struct {
		name string 
		age int
	}
	mystudent.name = "Nama"
	fmt.Println(mystudent.name)

	var stillstudent = struct {
		grade int
	}{39}
	fmt.Println(stillstudent.grade)
}