package main

import "fmt"

func main(){
	// example 1
	var names [5]string

	names[0] = "hai"
	names[1] = "bro"

	fmt.Println(names)
	fmt.Println(len(names))

	// example 2
	var myname = [2]string{"a","b"}
	fmt.Println(myname)

	// example 3
	var mynames = [2]string{
		"a",
		"b",
	}
	fmt.Println(mynames)

	// example 4
	var numbers = [...]int{2, 3, 2, 4, 3}
	fmt.Println(numbers)

	// example 5
	var dimension = [2] [3] int {
		[3] int {1,2,3},
		[3] int {4,5,6},
	}
	fmt.Println(dimension)

	for i:=0;i<len(dimension);i++ {
		fmt.Println(dimension[i])
	}

	for i,item := range dimension {
		fmt.Println(i)
		fmt.Println(item)
	}

	for _,item := range dimension {
		fmt.Println(item)
	}

	// example 6
	var fruits = make([]string, 2)

	fruits[0] = "apple"
	fruits[1] = "manggo"

	fmt.Println(fruits) 
}