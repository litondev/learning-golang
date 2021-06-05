package main

import "fmt"

type student struct {
	name string
	height float64
	age int32
	isGraduated bool
	hobbies []string
}

func main(){
	var data = student{
		name: "wick",
		height: 182.5,
		age: 26,
		isGraduated: true,
		hobbies: []string{"eating","slepping"},
	}

	// integer ke binner 
	fmt.Printf("%b \n",data.age)

	// integer ke karakter
	fmt.Printf("%c \n", 1235)

	// integer ke integer
	fmt.Printf("%d \n",data.age)

	// desimal ke notasi setandar
	fmt.Printf("%e \n",data.height)
	fmt.Printf("%E \n",data.height)

	// desimal ke desimal
	fmt.Printf("%f \n",data.height)
	fmt.Printf("%.2f \n",data.height)

	// desimal ke desimal tapi tidak bisa di custom nolnya
	fmt.Printf("%g \n",data.height)

	// integer ke oktal
	fmt.Printf("%o \n",data.age)

	// membuat data pointer
	fmt.Printf("%p \n",&data.name)

	// escape string
	fmt.Printf("%q \n", `" name \ height "`)

	// string
	fmt.Printf("%s\n", data.name)

	// boolean
	fmt.Printf("%t\n", data.isGraduated)

	// mengambil type variabel
	fmt.Printf("%T\n", data.name)

	// mengambil format apasaja
	fmt.Printf("%v\n", data)

	// mengambil struct dan propertinya
	fmt.Printf("%+v\n", data)

	// mengambil struct dan propertinya serta bagaiman di dekelarasikan
	fmt.Printf("%#v\n", data)

	// integer ke heksa
	fmt.Printf("%x\n", data.age)

	// menampilkan persent
	fmt.Printf("%%\n")
}