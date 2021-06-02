package main

import (
	"fmt"
)

func main(){
	var secret interface{}

	secret = "hai"
	fmt.Println(secret)


	secret = []string{"apple","buah"}
	fmt.Println(secret)

	var data map[string]interface{}
	data = map[string]interface{}{
		"name":      "ethan hunt",
		"grade":     2,
		"breakfast": []string{"apple", "manggo", "banana"},
	}
}