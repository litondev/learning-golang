package main

import "fmt"
import "net/http"
import "html/template"

func main(){
	http.HandleFunc("/",func(w http.ResponseWriter,r *http.Request){
		var data = map[string]string{
			"Name" : "jhon wick",
			"Message" : "have nice day",
		}

		var t,err = template.ParseFiles("web2.html")
		if err != nil {
			fmt.Println(err.Error())
		}

		t.Execute(w,data)
	})

	fmt.Println("Web Server On")
	http.ListenAndServe(":8080",nil)
}