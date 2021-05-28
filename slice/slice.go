package main

import "fmt"

func main(){
	myslice := [] string {"a","b"}
	fmt.Println(myslice)

	myslice1 := [] string {"a","b","c","d"}
	mynewslice1 := myslice1[2:3]
	fmt.Println(myslice1)
	fmt.Println(mynewslice1)



	var fruits = [] string {"apple", "grape", "banana", "melon"}

	var aFruits = fruits[0:3]
	var bFruits = fruits[1:4]

	var aaFruits = aFruits[1:2]
	var baFruits = bFruits[0:1]

	fmt.Println(fruits)   
	// [apple grape banana melon]fmt.Println(aFruits)  
	// [apple grape banana]fmt.Println(bFruits)  
	// [grape banana melon]fmt.Println(aaFruits) 
	// [grape]fmt.Println(baFruits) 
	// [grape]
	// Buah "grape" diubah menjadi "pinnaple"
	baFruits[0] = "pinnaple"
	fmt.Println(fruits)   // [apple pinnaple banana melon]
	fmt.Println(aFruits)  // [apple pinnaple banana]
	fmt.Println(bFruits)  // [pinnaple banana melon]
	fmt.Println(aaFruits) // [pinnaple]
	fmt.Println(baFruits) // [pinnaple]

	var fruits = []string{"apple", "grape", "banana"}
	var cFruits = append(fruits, "papaya")
	fmt.Println(fruits) 
	fmt.Println(cFruits) 


	var fruits = []string{"apple"}
	var aFruits = []string{"watermelon", "pinnaple"}
	var copiedElemen = copy(fruits, aFruits)
	fmt.Println(fruits)       

	// ["apple", "watermelon", "pinnaple"]
	fmt.Println(aFruits)      
	// ["watermelon", "pinnaple"]
	fmt.Println(copiedElemen) // 


	var fruits = []string{"apple", "grape", "banana"}var aFruits = fruits[0:2]var bFruits = fruits[0:2:2]fmt.Println(fruits)      // ["apple"
}