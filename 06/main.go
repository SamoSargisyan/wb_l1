package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

//Все способы останови выполнения горутины
// Контекст:
// 1.WithTimeOut
// 2.WithDeadline
// 3.WithCancel

func routineWithCancel(ctx context.Context, wg *sync.WaitGroup) {
	for i := 0; ; i++ {
		select {
		case <-ctx.Done():
			fmt.Println("I am killed with cancel!")
			wg.Done()
			return
		default:
			fmt.Println(i)
			time.Sleep(time.Second)
		}
	}
}

func routineWithTimeOut(ctx context.Context, wg *sync.WaitGroup) {
	for i := 0; ; i++ {
		select {
		case <-ctx.Done():
			fmt.Println("I am killed with timeOut!")
			wg.Done()
			return
		default:
			fmt.Println(i)
			time.Sleep(time.Second)
		}
	}
}
func routineWithDeadline(ctx context.Context, wg *sync.WaitGroup) {
	for i := 0; i >= 0; i++ {
		select {
		case <-ctx.Done():
			fmt.Println("I am killed with Deadline!")
			wg.Done()
			return
		default:
			fmt.Println(i)
			time.Sleep(time.Second)
		}
	}
}

func routineWithCraftedCancelChan(cancel chan bool, wg *sync.WaitGroup) {
	for i := 0; i >= 0; i++ {
		select {
		case <-cancel:
			fmt.Println("I am killed with CraftedCancelChannel!!")
			wg.Done()
			return
		default:
			fmt.Println(i)
			time.Sleep(time.Second)
		}
	}
}

func main() {
	var wg sync.WaitGroup
	cancelChan := make(chan bool) // handMade chan to stop Routine
	wg.Add(4)
	ctx1, cancel1 := context.WithCancel(context.Background())                             // when call cancel1() there's something in <-ctx.Done
	ctx2, _ := context.WithTimeout(context.Background(), 2*time.Second)                   // gets n Seconds until call cancel()
	ctx3, _ := context.WithDeadline(context.Background(), time.Now().Add(10*time.Second)) // gets timestamp when cancel() will be called
	go routineWithCancel(ctx1, &wg)
	go routineWithDeadline(ctx3, &wg)
	go routineWithTimeOut(ctx2, &wg)
	go routineWithCraftedCancelChan(cancelChan, &wg)
	go func() {
		time.Sleep(6 * time.Second) // ctx cancel management
		cancel1()
		time.Sleep(1 * time.Second)
		cancelChan <- true // handmade chan filled => signal to stop
	}()
	wg.Wait()
}
