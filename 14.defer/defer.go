package main

import "fmt"

func main(){
	defer fmt.Println("Hai Am Out")
	fmt.Println("Hello")
}