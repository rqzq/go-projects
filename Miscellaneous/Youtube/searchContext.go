package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	minLatency =10
	maxLatency = 5000
)

func main(){
	fmt.Println("Starting to Search for nyc-london")
	res,_:= search("nyc","London")
	fmt.Println("Got Results",res)
}

func search(from,to string)([]string,error){
	rand.Seed(time.Now().Unix())
	Latency:= rand.Intn(maxLatency-minLatency+1)-minLatency
	time.Sleep(time.Duration(Latency)*time.Millisecond)
	fmt.Printf("Started to search for %s-%s takes %dms...",from,to,Latency)
	return []string{from + "-" + to + "-british airways-11am", from + "-" + "delta airlines-12am"},nil
}

