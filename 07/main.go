/*
Реализовать конкурентную запись данных в map
*/
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

//go build -race
func touchMap(mux *sync.Mutex, wg *sync.WaitGroup, myMap map[int]int, key *int) {
	defer wg.Done()
	mux.Lock()
	myMap[*key] = rand.Intn(100)
	*key += 1
	fmt.Println(*key)
	mux.Unlock()
	time.Sleep(time.Second * 1)
}

func main() {
	var mux sync.Mutex
	var wg sync.WaitGroup
	myMap := make(map[int]int)
	var counter int
	nWorkers := 100
	for i := 0; i < nWorkers; i++ {
		wg.Add(1)
		go touchMap(&mux, &wg, myMap, &counter)
	}
	wg.Wait()
	fmt.Println("Everything OK? ", nWorkers == len(myMap)) //simple check
}
