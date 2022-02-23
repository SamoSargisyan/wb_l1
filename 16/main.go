package main

import (
	"fmt"
)

// https://qvault.io/golang/quick-sort-golang/

func quicksort(a []int, left, right int) {
	if right-left > 0 {
		_, pivotInd := partition(a, left, right)
		quicksort(a, left, pivotInd-1)
		quicksort(a, pivotInd+1, right)
	}
}

func partition(arr []int, low, high int) ([]int, int) {
	pivot := arr[high] // крайний элемент, конец массива
	i := low           // граничный элемент, который будет двигаться слева направо пошагово
	for j := low; j < high; j++ {
		if arr[j] < pivot { // если элемент массива меньше крайнего элемента, то кладем его слева от граничного элемента
			arr[j], arr[i] = arr[i], arr[j] // хоп
			i++                             // двигаем граничный элемент
		}
	}
	arr[i], arr[high] = arr[high], arr[i] // все элементы, больше кранего элемента - ведем вправо
	return arr, i
}

func main() {
	a := []int{5, 2, 7, 1, 4, 9, 3, 6, 8}

	quicksort(a, 0, len(a)-1)
	fmt.Println(a)
}
