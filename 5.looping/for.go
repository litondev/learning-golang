package main 

import "fmt"

func main(){

	for i:=0; i<5; i++ {
		fmt.Println(i)
	}

	p := 5
	for p < 10 {
		fmt.Println("Angka",p)
		p++
	} 

	s := 1
	for {
		fmt.Println("Angka",s)

		s++

		if s == 3 {
			continue		
		}

		if s == 5 {
			break
		}
	}
}