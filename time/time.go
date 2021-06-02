package main

import "fmt"
import "time"

func main(){
	var time1 = time.Now()
	fmt.Printf("time1 %v\n",time1)

	var time2 = time.Date(2011, 12, 24, 10, 20, 0, 0, time.UTC)    
	fmt.Printf("time2 %v\n", time2)

	fmt.Printf("Year %v\n",time1.Year())
}