package main

import "fmt"
import "encoding/json"

func main(){
	var jsonData = []byte(`{"Name": "Liton","Age": 1}`)

	var data interface{}

	var err = json.Unmarshal(jsonData,&data)
	if err != nil {
		fmt.Println(err.Error())
	}

	var decodeData = data.(map[string]interface{})
	fmt.Println("name : ",decodeData["Name"])
	fmt.Println("age : ",decodeData["Age"])

}