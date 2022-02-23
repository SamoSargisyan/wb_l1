package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type Counter struct {
	order int
}

func incrementCounter(ctx context.Context, mux *sync.Mutex, wg *sync.WaitGroup, counter *Counter) {
	for {
		select {
		case <-ctx.Done():
			defer wg.Done()
			return
		default:
			mux.Lock()
			counter.order += 1
			mux.Unlock()
			time.Sleep(time.Millisecond * 5000)
		}
	}
}

func main() {
	var wg sync.WaitGroup
	var mux sync.Mutex
	ctx, _ := context.WithTimeout(context.Background(), time.Second*10)

	counter := Counter{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go incrementCounter(ctx, &mux, &wg, &counter)
	}
	wg.Wait() // blocks until counter is not 0
	fmt.Println("Counted up to :", counter.order)
}
