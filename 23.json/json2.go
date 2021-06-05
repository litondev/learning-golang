package main

import "fmt"
import "encoding/json"


func main(){
	var jsonData = []byte(`{"Name": "Test","Age": 25}`)

	var data map[string]interface{}

	var err = json.Unmarshal(jsonData,&data)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println("user :",data["Name"])
	fmt.Println("age :",data["Age"])
}