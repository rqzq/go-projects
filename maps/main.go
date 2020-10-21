package main

import "fmt"

func main(){
	colors:= map[string]string{
		"red": "#ff0000",
		"white": "#ff986",
	}
	colors["orange"]="#ffffff"

	// Create Empty Map
	colo:= make(map[string]string)
	fmt.Println(colors)
	fmt.Println(colo)
}
