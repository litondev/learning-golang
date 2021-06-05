package main

import "fmt"
import "regexp"

func main(){
	var text = "testaA testbB testcC"
	var regex,_ = regexp.Compile(`[a-z]+`)

	var str1 = regex.FindAllString(text,-1)
	fmt.Println(str1)

	var str2 = regex.FindAllString(text,1)
	fmt.Println(str2)
}