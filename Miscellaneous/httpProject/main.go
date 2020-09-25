package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main(){
	response, err:= http.Get("https://api.coinbase.com/v2/prices/spot?currency=USD")
	if err != nil{
		fmt.Printf("The HTTP Request Failed with error %s \n", err)
	} else{
		data, _:= ioutil.ReadAll(response.Body)
		fmt.Println(string(data))
	}
	// Posting Json Data to API Endpoint
	jsonData:= map[string]string{
		"FirstName": "Bishwajit",
		"LastName": "Samanta",
	}
	jsonValue,_:= json.Marshal(jsonData)
	response, err = http.Post("https://httpbin.org/post","application/json",bytes.NewBuffer(jsonValue))
	if err != nil{
		fmt.Printf("The HTTP Request Failed with error %s \n", err)
	} else{
		data, _:= ioutil.ReadAll(response.Body)
		fmt.Println(string(data))
	}
}
