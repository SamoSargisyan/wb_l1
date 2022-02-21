package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func sender(ctx context.Context, wg *sync.WaitGroup, putHere chan int) {
	for i := 0; i >= 0; i++ {
		select {
		case <-ctx.Done():
			fmt.Println("I am sender and Im done")
			wg.Done()
			return
		default:
			putHere <- i
			time.Sleep(time.Second * 1)
		}
	}
}

func reader(ctx context.Context, wg *sync.WaitGroup, getHere chan int) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("I am reader and Im done by context")
			wg.Done()
			return
		case v := <-getHere:
			fmt.Println(v)
		}
	}
}

func main() {
	tube := make(chan int)
	nSecond := 0
	var wg sync.WaitGroup
	fmt.Println("How many second to work?")
	fmt.Scan(&nSecond)
	ctx, _ := context.WithTimeout(context.Background(), time.Second*time.Duration(nSecond))
	wg.Add(2)
	go reader(ctx, &wg, tube)
	go sender(ctx, &wg, tube)
	wg.Wait()
}
