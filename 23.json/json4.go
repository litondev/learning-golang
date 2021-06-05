package main

import "fmt"
import "encoding/json"

type User struct {
	Name string
}

func main(){
	var jsonString = `[
		{"Name": "Liton"},
		{"Name": "Roni"}
	]`

	var data []*User

	var err = json.Unmarshal([]byte(jsonString),&data)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println("name 1 :",data[0].Name)
	fmt.Println("name 2 :",data[1].Name)
}