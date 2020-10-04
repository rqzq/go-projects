package main

import "fmt"

var go_groceries []string

func add_groceries(a string){
	go_groceries=append(go_groceries,a)
}

func listt_groceries(){
	for i:=0;i<len(go_groceries);i++{
		fmt.Println("Below are the Groceries Added:",go_groceries[i])
	}
}

func main(){
	add_grocery("Bread")
	listt_groceries()

}
