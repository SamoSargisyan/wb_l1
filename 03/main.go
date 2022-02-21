package main

import (
	"fmt"
)

func main() {
	data := []int{2, 4, 6, 8, 10}
	sum := calculateSum(data)

	fmt.Printf("res = %d\n", sum)
}

func calculateSum(data []int) int {
	if len(data) < 1 {
		return 0
	}

	ch := make(chan int, len(data)) // инициализируем

	for i := 0; i < len(data); i++ {
		go posting(ch, data[i]) // вызываем
	}

	var sum int
	for i := 0; i < len(data); i++ {
		sum += <-ch // суммируем все, шо транслировалось в канал
	}

	return sum // возвращаем
}

func posting(ch chan int, value int) {
	ch <- value * value // пишем в канал
}