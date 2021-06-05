package main

import "fmt"
import "strconv"

func main(){
	var str = "124"
	var num,err = strconv.ParseInt(str,10,64)
	if err == nil {
		fmt.Println(num)
	}

	str = "1010"
	num,err = strconv.ParseInt(str,2,8)
	if err == nil {
		fmt.Println(num)
	}
}