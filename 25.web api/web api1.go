package main

import (
	"fmt"
	"net/http"
	"encoding/json"
)

type Student struct {
	Id string
	Name string
	Grade int
}

var data = []Student {
	{Id: "1",Name: "testing1",Grade: 10},
	{Id: "2",Name: "testing2",Grade: 20},
}

func users(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/json")

	if r.Method != "GET" {
		http.Error(w,"",http.StatusBadRequest)
		return 
	}

	var result,err = json.Marshal(data)

	if err != nil {
		http.Error(w,err.Error(),http.StatusInternalServerError)
		return 
	}

	w.Write(result)
}

func user(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/json")

	if r.Method != "GET" {
		http.Error(w,"",http.StatusBadRequest)
	}

	var id = "1"
	var result []byte
	var err error

	for _,each := range data {
		if each.Id == id {
			result,err = json.Marshal(each)

			if err != nil {
				http.Error(w,err.Error(),http.StatusInternalServerError)
				return 
			}

			w.Write(result)
			return 
		}
	}

	http.Error(w,"User not found",http.StatusBadRequest)
}

func main(){
	http.HandleFunc("/user",user)
	http.HandleFunc("/users",users)

	fmt.Println("Server Running")
	http.ListenAndServe(":8000",nil)
}