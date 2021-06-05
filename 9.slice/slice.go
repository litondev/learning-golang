package main

import "fmt"

func main(){
	// example 1
	myslice := [] string {"a","b"}
	fmt.Println(myslice)

	// example 2
	myslice1 := [] string {"a","b","c","d"}
	mynewslice1 := myslice1[2:3]
	fmt.Println(myslice1)
	fmt.Println(mynewslice1)

	// example 3
	var fruits = [] string {"apple", "grape", "banana", "melon"}
	
	var aFruits = fruits[0:3]
	var bFruits = fruits[1:4]

	var aaFruits = aFruits[1:2]
	var baFruits = bFruits[0:1]

	fmt.Println(fruits)   
	
	baFruits[0] = "pinnaple"

	fmt.Println(fruits)  
	fmt.Println(aFruits) 
	fmt.Println(bFruits)  
	fmt.Println(aaFruits) 
	fmt.Println(baFruits) 

	// example 4
	var fruits = []string{"apple", "grape", "banana"}
	var cFruits = append(fruits, "papaya")
	fmt.Println(fruits) 
	fmt.Println(cFruits) 

	// example 5
	var fruits = []string{"apple"}
	var aFruits = []string{"watermelon", "pinnaple"}
	var copiedElemen = copy(fruits, aFruits)
	fmt.Println(fruits)       
	fmt.Println(aFruits)      
	fmt.Println(copiedElemen) 

	// example 6
	var fruits = []string{"apple", "grape", "banana"}
	var aFruits = fruits[0:2]
	var bFruits = fruits[0:2:2]
	fmt.Println(fruits)   
}