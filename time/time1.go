package main 

import "fmt"
import "time"

func main(){
	var layoutFormat,value string
	var date time.Time

	layoutFormat = "2006-01-02 15:04:05"
	value = "2015-09-02 08:04:00"
	date, _ = time.Parse(layoutFormat, value)
	fmt.Println(value, "\t->", date.String())	

	layoutFormat = "02/01/2006 MST"
	value = "02/09/2015 WIB"
	date, _ = time.Parse(layoutFormat, value)
	fmt.Println(value, "\t\t->", date.String())

	var dateS1 = date.Format("Monday 02, January 2006 15:04 MST")
	fmt.Println("dateS1", dateS1)

	var date, err = time.Parse("06 Jan 15", "02 Sep 15 08:00 WIB")
	
	if err != nil {    
		fmt.Println("error", err.Error())
		return
	}

	fmt.Println(date)
}