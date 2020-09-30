package main

import (
	"fmt"
	"gopkg.in/src-d/go-git.v4"
	"os"
	"os/exec"
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
	// Installing Docker
	command:= exec.Command("/bin/sh","-c","sudo apt-get update && apt-get install  apt-transport-https ca-certificates  curl gnupg-agent software-properties-common ")
	if err:= command.Run(); err!=nil{
		fmt.Println("Error in Installation", err)
	}
	
}