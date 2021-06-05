package main

import "struct11/library"

func main(){
	// can access
	library.SayHello()

	// cannot access
	// library.intro()
	library.Hai()
}