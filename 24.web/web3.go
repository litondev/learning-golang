package main

import "fmt"
import "net/url"

func main(){
	var urlString = "http://depeloper.com:80/hello?name=john wick&age=27"
	var u,e = url.Parse(urlString)

	if e != nil {
		fmt.Println(e.Error())
	}	

	fmt.Printf("Url : %v \n",urlString)
	fmt.Printf("Schema : %v \n",u.Scheme)
	fmt.Printf("Host : %v \n",u.Host)
	fmt.Printf("Path : %v \n",u.Path)

	var name = u.Query()["name"][0]
	var age = u.Query()["age"][0]
	fmt.Printf("name: %v , age %v \b",name,age)
}