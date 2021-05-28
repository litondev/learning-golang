package main

import "fmt"

func main(){
	var allStudent = []struct{
		name string
		age int
	}{
		{name : "h"},
		{age : 4},
		{name : "a"},
	}

	for _,e := range allStudent{
		fmt.Println(e)
	}
}