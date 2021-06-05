package main

import "fmt"
import "regexp"

func main(){
	var text = "bananan testing"
	var regex,_ = regexp.Compile(`[a-z]+`)

	var idx = regex.FindStringIndex(text)
	fmt.Println(idx)

	var str = text[idx[0]:idx[1]]
	fmt.Println(str)
}