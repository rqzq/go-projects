package main

import "fmt"

var go_groceries []string

func add_groceries(a string){
	go_groceries=append(go_groceries,a)
}

func listt_groceries(){
	fmt.Println("Groceries List are as follows: ")
	for _,data:= range go_groceries{
		fmt.Println(data)
	}
}

func main(){
	add_groceries("Bread")
	add_groceries("Wheat")
	add_groceries("Rice")
	add_groceries("Omlette")
	add_groceries("Ice Cream")
	listt_groceries()

}
