package main

import (
	"fmt"
	"sync"
)

func main() {
	inputArr := []int{2, 4, 6, 8, 10}
	wg := new(sync.WaitGroup)

	for _, i := range inputArr {
		wg.Add(1)
		go pow(wg, i)
	}

	wg.Wait()
}

func pow(wg *sync.WaitGroup, i int) {
	defer wg.Done()
	fmt.Println(i * i)
}

// то же самое но в closer
func secondMethod() {
	wg := new(sync.WaitGroup) // вроде как ждет, что бы все гоурутины закончили выполнение, и тогда получим результат

	nums := []int{2, 4, 6, 8, 10}
	for _, i := range nums {
		wg.Add(1)        // добавляем в счетчик по 01 гоурутине каждую итерацию цикла
		go func(a int) { // это значение берется из пендинга на 3 сточки ниже
			defer wg.Done() // сверху прибавили, а тут убавляем
			fmt.Println(a * a)
		}(i) // отсюда
	}

	wg.Wait() // когда цикл закончит отрабатывать, число гоурутин в счетчике будет 0. Этого и ждем
}
