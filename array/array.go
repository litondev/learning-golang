package main

import "fmt"


func main(){
	var names [5]string
	names[0] = "hai"
	names[1] = "bro"

	fmt.Println(names)
	fmt.Println(len(names))

	var myname = [2]string{"a","b"}
	fmt.Println(myname)

	var mynames = [2]string{
		"a",
		"b",
	}
	fmt.Println(mynames)

	var numbers = [...]int{2, 3, 2, 4, 3}
	fmt.Println(numbers)

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

	var fruits = make([]string, 2)

	fruits[0] = "apple"
	fruits[1] = "manggo"

	fmt.Println(fruits) 
}