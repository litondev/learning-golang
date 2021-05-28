package main

import (
	"fmt"
	"strings"
)

// type FilterCallback func(string) bool

// func filter(data []string, callback FilterCallback) []string {

func filter(data []string,callback func(string) bool) []string {
	var result []string 

	for _,e := range data {
		if filtered := callback(each); filtered {
			result = append(result,each)
		}
	}

	return result
}

func main(){
	var data = []string{"wick","a"}
	var dataContainsO = filter(data,func(each string) bool{
			return strings.Contains(each,'a')
		})
	fmt.Println(dataContainsO)
}

