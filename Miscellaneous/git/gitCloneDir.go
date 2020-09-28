package main

import (
	"fmt"
	"gopkg.in/src-d/go-git.v4"
	"os"
)

func main(){
	// Get Current Working Directory
	mydir, err1:= os.Getwd()
	if err1 != nil {
		fmt.Println(err1)
	}
	// Cloning Github Repo into Current Working Directory
	_,err:=git.PlainClone(mydir,false,&git.CloneOptions{URL: "https://github.com/RealOrko/concourse-fetch.git",Progress: os.Stdout})
	if err != nil {
		fmt.Println(err)
	}
}