package main

import (
	"fmt"
	"gopkg.in/src-d/go-billy.v4/memfs"
	git "gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/storage/memory"
	"io"
	"os"
)

func main(){
	fs := memfs.New()
	_, err := git.Clone(memory.NewStorage(),fs, &git.CloneOptions{
		URL: "https://github.com/RealOrko/concourse-fetch.git",
		Progress: os.Stdout,
	})
	if err != nil{
		panic(err)
	}
	readme, err := fs.Open("README.md")
	if err != nil{
		fmt.Println(err)
	}
	io.Copy(os.Stdout,readme)
}
