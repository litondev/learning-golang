package main

import "fmt"
import "regexp"

func main(){
	var text = "testa testb testc"
	var regex,_ = regexp.Compile(`[a-z]+`)

	var str1 = regex.ReplaceAllString(text,"hai")
	fmt.Println(str1)
}