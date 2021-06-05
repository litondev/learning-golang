package main

import "fmt"
import "strconv"

func main(){
	var str = "24.24"
	var num,err = strconv.ParseFloat(str,32)
	if err == nil {
		fmt.Println(num)
	}
}