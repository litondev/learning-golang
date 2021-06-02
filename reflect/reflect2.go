package main 

import (
	"fmt"
	"reflect"
)

func main(){
	var number = 23
	var reflectValue = reflect.ValueOf(number)

	fmt.Println("tipe variabel : ",reflectValue.Type())
	fmt.Println("nilai variabel : ",reflectValue.Interface())
	// var nilai = reflectValue.Interface().(int)
}