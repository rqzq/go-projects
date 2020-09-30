package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main(){

	//Collect Environment Variables passed from Params Section in Concourse
	appName:= os.Getenv("app")
	environment:= os.Getenv("env")
	tenant:= os.Getenv("tenant")
	resp, err:= http.Get("https://"+ tenant + appName +"-app-" + environment +".apps.prod.smarsh.cloud/actuator/health")

	//Error Handling Section
	if err != nil{
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// Application Working Condition based on the Status Code
	if resp.StatusCode == 200 {
		fmt.Println("Application",appName, "is working on",environment,"environment!")
	} else {
		fmt.Println("Application",appName, "is Broken on",environment,"environment!")
	}
}