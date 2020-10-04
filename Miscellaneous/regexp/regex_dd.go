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

	var strArray [10]string

	// Getting the Values from Params Section
	text:= os.Getenv("TEXT")
	title:= os.Getenv("TITLE")
	dd_key:= os.Getenv("DD_CLIENT_API_KEY")

	// Setting Array for Tags Section

	strArray[1] = fmt.Sprintf("Environment: %s",os.Getenv("ENV_NAME"))
	strArray[2] = fmt.Sprintf("Customer: %s",os.Getenv("CUSTOMER"))
	strArray[3] = fmt.Sprintf("Cloud: %s",os.Getenv("CLOUD"))
	strArray[4] = fmt.Sprintf("Product: %s",os.Getenv("PRODUCT"))
	strArray[5] = fmt.Sprintf("Version: %s",os.Getenv("VERSION"))
	strArray[6] = fmt.Sprintf("Event Type: %s",os.Getenv("EVENT_TYPE"))

	// Creating Struct for Querying Specific Fields
	type Datadog struct {
		Status string `json:"status"`
		Event struct{
			URL string `json:"url"`
		} `json:"event"`
	}
	var event Datadog

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
	jsonData:= map[string]interface{}{
		"text": text,
		"title": title,
		"source_type_name": "API",
		"tag": strArray,
	}
	jsonValue,_:= json.Marshal(jsonData)
	fmt.Println(string(jsonValue))
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