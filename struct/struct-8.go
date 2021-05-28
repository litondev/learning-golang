package main

import "fmt"

type person struct {
	name string `hai`
}

func main() {
	var myPerson = person{}
	fmt.Println(myPerson.name)
}