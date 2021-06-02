package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

func sendData(ch chan<- int){
	for i := 0;true;i++ {
		ch <- i
		time.Sleep(time.Duration(rand.Int()%10+1) * time.Second)
	}
}

func retreiveData(ch <-chan int) {    
	loop:
	for {
		select {
			case data := <-ch:            
				fmt.Print(`receive data "`, data, `"`, "\n")
			case <-time.After(time.Second * 5):            
				fmt.Println("timeout. no activities under 5 seconds")
				break loop        
		}    
	}
}


func main(){
 	rand.Seed(time.Now().Unix())    
 	runtime.GOMAXPROCS(2)
 	var messages = make(chanint)
 	go sendData(messages)    
 	retreiveData(messages)
}