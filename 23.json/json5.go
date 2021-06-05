package main

import "fmt"
import "encoding/json"

type User struct {
	Name string
}

func main(){
	var object = []User{
		{Name: "Soni"},
		{Name: "Ronit"},
	}

	var jsonData,err = json.Marshal(object)
	if err != nil {
		fmt.Println(err.Error())
	}

	var jsonString = string(jsonData)
	fmt.Println(jsonString)
}