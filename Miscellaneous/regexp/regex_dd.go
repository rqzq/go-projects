package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
)

func main(){

	// Reading the File Content
	content, err:= ioutil.ReadFile("version.txt")
	if err != nil{
		fmt.Println(err)
	}
	version:= string(content)

	// Applying the Regular Expression for matching the Characters
	re:= regexp.MustCompile(`(?m)[\d\.]+[\.-][\d]{2,}`)
	for _, match:= range(re.FindAllString(version,-1)){
		fmt.Println(match)
	}

	//Executing Curl Statements for Posting Datadog Events
}
