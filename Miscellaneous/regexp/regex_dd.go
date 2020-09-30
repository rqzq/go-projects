package main

import (
	"fmt"
	"regexp"
)

func main(){
	var version = "storm-0.0.197.tgz"
	re:= regexp.MustCompile(`(\d+).(\d+).(\d+)`)
	result := re.FindAllStringSubmatch(version,-1)
	for res := range(result){
		fmt.Printf("Version:%s",result[res][0])
	}
	var appVersion = "restapi-app-3.3.8.2-3.3.5-3811-g154be7c-1597293962429.zip"
	regex:= regexp.MustCompile(`-(\d+).(\d+).(\d+)-(\d+)`)
	result1 := regex.FindAllStringSubmatch(appVersion,-1)
	for res := range(result1){
		fmt.Printf("Version:%s",result1[res][0])
	}
}
