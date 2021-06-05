package main

import "fmt"
import "regexp"

func main(){
	var text = "bannana burger soup"
	var regex,err = regexp.Compile(`[a-z]+`)

	if err != nil {
		fmt.Println(err.Error())
	}

	var res1 = regex.FindAllString(text,2)
	fmt.Println("%#v \n",res1)

	var res2 = regex.FindAllString(text,-1)
	fmt.Println("%#v \n",res2)
}