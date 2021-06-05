package main

import "fmt"

func main(){
	isTrue := true 
	
	isNotTrue := false

	if isTrue {
		fmt.Println("Is True")
	}

	if isNotTrue {
		fmt.Println("Is Not True")
	} else if !isNotTrue {
		fmt.Println("Is Not True False")
	} else {
		fmt.Println("Else")
	}
}