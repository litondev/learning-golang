package main

import "fmt"
import "encoding/json"

type User struct {
	Name string
	Age int
}

func main(){
	var jsonString = `{"Name": "Jhon Wick", "Age": 27}`
	var jsonData = []byte(jsonString)

	var data User

	var err = json.Unmarshal(jsonData,&data)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println("user :",data.Name)
	fmt.Println("age :",data.Age)
}