package main

import (
	"fmt"
	"regexp"
)

func main(){
	var sentence = "Username: Bishwajit\nUsername:Samanta"
	parts:= regexp.MustCompile(":\\s*").Split(sentence,-1)
	for _,part:= range(parts){
		fmt.Println(part)
	}
}
