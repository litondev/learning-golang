package main

import "fmt"

func main(){
	var allStudent = []struct{
		name string
		age int
	}{
		{name : "Nama"},
		{age : 4},
		{name : "Nama"},
	}

	for _,e := range allStudent{
		fmt.Println(e)
	}
}