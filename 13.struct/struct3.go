package main 

import "fmt"

func main(){
	var studentStruct = struct {
		name string
		age int
	}{name:"Nama"}

	fmt.Println(studentStruct.name)
}