package main

import "fmt"

type Person struct {
	firstName string
	lastName string
}
func main(){
	alex:= Person{
		firstName: "Bishwajit",
		lastName:  "Samanta",
	}
	fmt.Println(alex.firstName)
}
