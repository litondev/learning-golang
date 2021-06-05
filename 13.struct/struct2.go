package main 

import "fmt"

type people struct {
	age int
	name string
}

type student struct {
	class int
	name string
	people 
}

func main(){
	var studentA = student{}
	studentA.people.name = "liton"
	studentA.age = 15
	fmt.Println(studentA.people.age)

	var peopleA = people{age:4,name:"asd"}
	var studentB = student{people:peopleA,name:"asd"}
	fmt.Println(studentB.name)
}