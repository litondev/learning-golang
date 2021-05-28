package main

import "fmt"
import "strings"

func main(){
	var names = []string{"John", "Wick"}
	printMessage("hello",names)
}

func printMessage(message string, arr []string) {
	var nameString = strings.Join(arr, " ")    
	fmt.Println(message, nameString)
}