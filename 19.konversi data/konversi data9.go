package main

import "fmt"

func main(){
	var data = map[string]interface{}{
		"nama" : "jhon wick",
	}
	fmt.Println(data["nama"].(string))
}