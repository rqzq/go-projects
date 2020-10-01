package main

import "fmt"

type Data struct {
	Status string `json:"status"`
	Event struct{
		ID int64 `json:"id"`
		Title string `json:"title"`
		Tags string `json:"tags"`
	} `json:"event"`
}

func main(){
	fmt.Println()
}
