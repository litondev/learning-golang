package main

import "fmt"

func main(){
	point := 3

	switch point {
	case 8:
		fmt.Println("8")
	case 6:
		fmt.Println("6")
	default:
		fmt.Println("default")
	}

	switch point {
		case 9:
			fmt.Println("9")
		case 8,7,6:
			fmt.Println("8,7,6")
		default:
			fmt.Println("default")
	}
	
	switch {
		case point >= 9:	
			fmt.Println("lebih besar dari pada 9")
		case (point >= 5) && (point <= 7):
			fmt.Println("lebih besar dari pada 5 dan lebih kecil dari pada 7")
			fallthrough
		default:{
			fmt.Println("default")
		}
	}
}