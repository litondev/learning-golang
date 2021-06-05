package main

import "fmt"

type person struct {
	name string `Nama`
}

func main() {
	var myPerson = person{}
	fmt.Println(myPerson.name)
}