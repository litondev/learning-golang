package main

import "fmt"
import "net/http"

func index(w http.ResponseWriter,r *http.Request){
	fmt.Fprintln(w,"apa kabar!")
}

func main(){
	http.HandleFunc("/",func(w http.ResponseWriter,r *http.Request){
		fmt.Fprintln(w,"Hello")
	})
	http.HandleFunc("/index",index)

	fmt.Println("starting web server")
	http.ListenAndServe(":8080",nil)
}