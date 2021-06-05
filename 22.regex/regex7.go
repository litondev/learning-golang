package main

import "fmt"
import "regexp"

func main(){
	var text = "haia haib haic"
	var regex,_ = regexp.Compile(`[a-z]+`)

	var str1 =  regex.ReplaceAllStringFunc(text,func(each string) string {
		if each == "haib" {
			return "super"
		}
		return each
	})

	fmt.Println(str1)
}