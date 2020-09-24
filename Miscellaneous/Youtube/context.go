package main

import (
	"context"
	"fmt"
	"time"
)

func main(){
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	time.AfterFunc(time.Second,cancel)
	sleepAndTalk(ctx, 5 * time.Second,"Hello")
}

func sleepAndTalk(ctx context.Context, d time.Duration, msg string){
	fmt.Println(msg)
}
