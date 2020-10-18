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
	jim.print()
	alex.print()
}

func (p Person) updateName(newFirstName string){
	p.firstName=newFirstName
}

func (p Person) print(){
	fmt.Println(p.firstName)
	fmt.Println(p.contact.email)
}