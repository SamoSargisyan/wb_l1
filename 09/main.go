package main

import "fmt"

func main() {
	array := []int{1, 2, 3, 10, 20, 30, 40, 100, 200, 300}
	for result := range doubleNumbersChanel(simpleChanel(array...)) {
		fmt.Println(result)
	}
}

func simpleChanel(numbers ...int) <-chan int {
	chanel := make(chan int)
	go func() {
		for _, number := range numbers {
			chanel <- number
		}
		close(chanel)
	}()
	return chanel
}

func doubleNumbersChanel(chanel <-chan int) <-chan int {
	x2chanel := make(chan int)
	go func() {
		for number := range chanel {
			x2chanel <- 2 * number
		}
		close(x2chanel)
	}()
	return x2chanel
}
