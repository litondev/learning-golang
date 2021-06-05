package main

import "fmt"
import "regexp"

func main(){
	var text = "bananan test aku"
	var regex,_ = regexp.Compile("[a-z]+")

	var str1 = regex.Split(text,-1)
	fmt.Println(str1)
}
