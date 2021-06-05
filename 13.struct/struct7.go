package main

import "fmt"

func main(){
	type person struct {name string;age int;}
	var myPerson = person{name : "Nama"}
	fmt.Println(myPerson.name)
}