package main

import (
	"fmt"
)

const g_cap int=3
var g_groceries[g_cap]string
var g_len =0
var count = 0

func add_grocery(a string){
	if g_len<g_cap{
		g_groceries[g_len]=a
		g_len++
	} else{
		count+=1
	}
}

func list_groceries(){
	fmt.Println("Grocery List are as follows:")
	for i:=0;i<g_len;i++{
		fmt.Println(g_groceries[i])
	}
	fmt.Printf("There are these Many Items Added Extra: %d \n", count)
}

func main(){
	add_grocery("Bread")
	add_grocery("Cucumber")
	add_grocery("Omlette")
	list_groceries()
}
