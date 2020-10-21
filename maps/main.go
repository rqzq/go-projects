package main

import "fmt"

func main(){
	colors:= map[string]string{
		"red": "#ff0000",
		"white": "#ff986",
		"orange": "#fffff",
	}
	printMap(colors)
}

func printMap(c map[string]string){
	for color, hex:= range c{
		fmt.Println("Hex of ", color, "is", hex)
	}
}