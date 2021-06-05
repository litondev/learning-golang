package main 

import "fmt"

// chan mengirim dan menerima data

// chan<- mengirim data 
func sendMessage(ch chan<- string) {
	for i := 0; i < 20; i++ {        
		ch <- fmt.Sprintf("data %d", i)    
	}

	close(ch)
}

// <-chan menerima data
func printMessage(ch <-chan string) {
	for message := range ch {
	    fmt.Println(message)    
	}
}

func main(){
	runtime.GOMAXPROCS(2)

	var messages = make(chanstring)
	go sendMessage(messages)    

	printMessage(messages)
}