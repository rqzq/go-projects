package main

import "fmt"

type ContactInfo struct {
	email string
	zipCode int
}

type Person struct {
	firstName string
	lastName string
	contact ContactInfo
}
func main(){
	alex:= Person{
		firstName: "Bishwajit",
		lastName:  "Samanta",
	}
	jim:= Person{
		firstName: "Payel",
		lastName:  "Samanta",
		contact:   ContactInfo{
			email:   "bishwajitsamanta1689@gmail.com",
			zipCode: 713205,
		},
	}
	fmt.Println(alex.firstName)
	fmt.Printf("%+v",jim)
	fmt.Println(jim.contact.email)
}
