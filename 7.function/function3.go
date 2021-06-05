package main

import "fmt"

func main(){
	myslice := []int{1,2,3,4}

	avg := hitung(myslice...)

	fmt.Println(avg)
}

func hitung(numbers ...int) int {
	mynumber := 0

	for _,item := range numbers {
		mynumber += item
	}

	return mynumber
}