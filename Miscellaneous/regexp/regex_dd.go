package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
)

func main(){
	content, err:= ioutil.ReadFile("version.txt")
	if err != nil{
		fmt.Println(err)
	}
	version:= string(content)
	re:= regexp.MustCompile(`(?m)[\d\.]+[\.-][\d]{2,}`)
	for _, match:= range(re.FindAllString(version,-1)){
		fmt.Println(match)
	}
}
