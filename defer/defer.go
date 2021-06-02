package main

import "fmt"

func main(){
	defer fmt.Println("Hai")
	fmt.Println("Hello")
}