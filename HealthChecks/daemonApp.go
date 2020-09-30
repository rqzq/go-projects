package main

import (
	"fmt"
	"log"
	"net/http"
)

func main(){
	resp, err:= http.Get("https://pstexport-daemon-app-dev-2.apps.prod.smarsh.cloud/actuator/health")
	if err != nil{
		log.Fatal(err)
	}
	defer resp.Body.Close()
	// Print the Status Code
	if resp.StatusCode == 200 {
		fmt.Println("Application is working!")
	} else {
		fmt.Println("Application is Broken!")
	}
}
