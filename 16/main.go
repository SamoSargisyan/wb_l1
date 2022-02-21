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
	pivot := arr[high]
	i := low
	for j := low; j < high; j++ {
		if arr[j] < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i], arr[high] = arr[high], arr[i]
	return arr, i
}

func main() {
	a := []int{5, 2, 7, 1, 4, 9, 3, 6, 11}

	quicksort(a, 0, len(a)-1)
	fmt.Println(a)
}
