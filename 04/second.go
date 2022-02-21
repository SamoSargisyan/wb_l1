package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"time"
)

func mainWriter(ctx context.Context, wg *sync.WaitGroup, intMun chan int) {
	for i := 0; i >= 0; i++ {
		select {
		case <-ctx.Done(): // when cancel() is called we finish goroutine
			fmt.Println("I am mainWriter and I am stopped with context")
			wg.Done()
			return
		default: // payload
			fmt.Printf("mainWriter puts: %d\n", i)
			intMun <- i
			time.Sleep(time.Millisecond * 1000)
		}
	}
}

func regularWorker(ctx context.Context, wg *sync.WaitGroup, intNum chan int, chanName int, ) {
	for i := 0; i >= 0; i++ {
		select {
		case <-ctx.Done(): // when cancel() is called we finish goroutine
			fmt.Println("I am regularWorker ", chanName, " And I've got stopped with context")
			wg.Done()
			return
		case v := <-intNum: //payload
			fmt.Printf("I AM WORKER№ %d AND I GIVE YOU FROM CHANNEL: %d\n", chanName, v)
			time.Sleep(time.Millisecond * 4000)
		}
	}
}

func main() {
	var (
		intNumbers chan int
		nWorkers   int
		wg         sync.WaitGroup
	)
	ctx, cancel := context.WithCancel(context.Background())
	// we get cancel fuction here
	// If we need to send a signal to stop smth then call cancel()
	//But why is it better than handMaded "stop" chan? Best practise?
	intNumbers = make(chan int)
	fmt.Println("How many workers?")
	fmt.Scan(&nWorkers)
	wg.Add(1)
	go mainWriter(ctx, &wg, intNumbers)
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt)
		<-c      // waiting for Ctrl+C
		cancel() // send signal for every goRoutine with ctx
	}()
	for i := 0; i < nWorkers; i++ {
		wg.Add(1)
		go regularWorker(ctx, &wg, intNumbers, i+1)
	}
	fmt.Println("wg.wait blocked")
	wg.Wait()
	fmt.Println("YOU PRESSED CTRL+C")
	/*
		What's the benefit of handling Ctrl+C?
		Maybe this way we let our goroutine to finish predictably
		(wait until payload of each goroutine is finished)
	*/
}