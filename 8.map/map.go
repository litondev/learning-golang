package main

import "fmt"

func main(){
	// example 1
	var chicken map[string]int 
	chicken = map[string]int{}

	chicken["januari"] = 50
	chicken["februari"] = 40

	fmt.Println("januari", chicken["januari"])
	fmt.Println("mei",     chicken["mei"])     

	// example 2 
	mychik := map[string]int{"a":1,"b":2}
	mychik1 := map[string]int{
		"a" : 1,
		"b" : 2,
	}

	fmt.Println(mychik)
	fmt.Println(mychik1)

	// example 3
	var mydate = map[string]int{"a":1,"b":2}
	delete(mydate, "a")
	for key,value := range mydate {
		fmt.Println(key)
		fmt.Println(value)
	}

	// example 4
	chicken = map[string]int{"a":1,"b":2}
	var value,isExist = chicken["c"]

	if isExist {
		fmt.Println(value)
	} else {
		fmt.Println("Tidak Ada")
	}
}