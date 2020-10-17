package main

import "fmt"

func makeRange(min,max int)[]int{
	a:= make([]int,max-min+1)
	for i:= range a{
		a[i] = min +i
	}
	return a
}

func main(){
	evenOdd:= makeRange(10,20)
	fmt.Println(evenOdd)

	for i, value:= range evenOdd{
		if evenOdd[i] %2 ==0{
			fmt.Printf("The number %d is even \n", value)
		} else {
			fmt.Printf("The number %d is odd \n", value)
		}
	}
}