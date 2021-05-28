package main

import "fmt"

func main(){
	numbers := []int{1,2,3,4,5}

	newNumbers := func(min int) []int {
		var r []int

		for _,e := range numbers {
			if e < min {
				continue
			}

			r = append(r,e)
		}

		return r
	}(3)

	fmt.Println("Print : ",numbers)
	fmt.Println("Print : ",newNumbers)
}