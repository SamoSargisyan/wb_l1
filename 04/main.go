// https://gobyexample.com/signals - полезная ссылка

package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// воркеры завершает работу когда закроется канал от куда он читает
func worker(ch <-chan interface{}, id int) {
	// читаем из канала пока он открытый
	for d := range ch {
		fmt.Printf("worker %d: %v\n", id, d)
	}
}

func main() {
	var workersNum int

	fmt.Print("Input number of workers: ")
	if _, err := fmt.Scan(&workersNum); err != nil {
		log.Fatalln("Bad number of workers value, program is shutting down")
	}
	// как я понял, гоурутины можно остановить только с помощью канала или контекста.
	// что бы отправить сигнал, нам нужен канал
	done := make(chan os.Signal, 1) //канал для системного сигнала
	data := make(chan interface{})  //канал типа interface{}, так как данные - произвольные

	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM) // Notify транслирует входящие сигналы в канал, который указан первым аргументом
	for i := 0; i < workersNum; i++ {
		go worker(data, i)
	}
	for {
		select {
		case <-done:
			close(data)
			close(done)
			return
		default:
			data <- rand.Intn(1000)
			time.Sleep(time.Millisecond * 400)
		}
	}
}
