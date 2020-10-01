package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
)

func main(){

	// Creating Struct for Querying Specific Fields
	type Datadog struct {
		Status string `json:"status"`
		Event struct{
			URL string `json:"url"`
		} `json:"event"`
	}
	var event Datadog

	// Getting the Values from Params Section
	text:= os.Getenv("TEXT")
	dd_key:= os.Getenv("DD_CLIENT_API_KEY")
	title:= os.Getenv("TITLE")

	// Reading the File Content
	content, err:= ioutil.ReadFile("version")
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
	jsonData:= map[string]string{
		"text": text,
		"title": title,
		"source_type_name": "API",
	}
	jsonValue,_:= json.Marshal(jsonData)
	request, _:= http.NewRequest("POST","https://api.datadoghq.com/api/v1/events",bytes.NewBuffer(jsonValue))
	request.Header.Set("Content-Type","application/json")
	request.Header.Set("DD-API-KEY",dd_key)
	client:= &http.Client{}
	response,err := client.Do(request)
	if err != nil{
		log.Fatalln(err)
	} else{
		data,_:= ioutil.ReadAll(response.Body)
		json.Unmarshal(data,&event)
		fmt.Println("Status Code:",event.Status)
		fmt.Println("Datadog Event URL:",event.Event.URL)
	}
}