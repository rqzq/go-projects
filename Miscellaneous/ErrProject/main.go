package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Website string `json:"website"`
}

func main(){
	bs:= Person{
		FirstName: "Bishwajit",
		LastName:  "Samanta",
		Website:   "smarsh.com",
	}
	data, _ := json.MarshalIndent(bs,""," ")
	fmt.Println(string(data))
	json.Unmarshal(data,&bs)
	fmt.Println(bs)

}
