package main

import (
	"fmt"
)


func main() {
	a,b := countThis(10,10)

	fmt.Println(a)
	fmt.Println(b)
}

func countThis(a,b int) (int,string) {
	return a,"b"
}