package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	go HandleRequest(ctx)
	for {
	}
}

func HandleRequest(ctx context.Context) {
	go WriteRedis(ctx)
	go WriteDataBase(ctx)
	for {
		select {
		case <-ctx.Done():
			fmt.Println("HandleRequest done..")
			return
		default:
			time.Sleep(time.Second)
			fmt.Println("HandleRequest is running..")
		}
	}
}

func WriteRedis(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("WriteRedis done..")
			return
		default:
			time.Sleep(time.Second)
			fmt.Println("WriteRedis is running..")
		}
	}
}

func WriteDataBase(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("WriteDataBase done..")
			return
		default:
			time.Sleep(time.Second)
			fmt.Println("WriteDataBase is running..")
		}
	}
}
