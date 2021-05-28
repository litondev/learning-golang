package main 

import "fmt"

func main(){
	var studentStruct = struct {
		name string
		age int
	}{name:"hai"}

	fmt.Println(studentStruct.name)
}