package main

import (
	"fmt"
)

func remove(slice []int, s int) []int {
	return append(slice[:s], slice[s+1:]...)
}

func removeNotOrdered(s []int, i int) []int {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7}
	remove(a, 3) // keeps order in collections but slower
	fmt.Println(a)
	a = []int{1, 2, 3, 4, 5, 6, 7}
	removeNotOrdered(a, 3) // doesn't keep order but faster
	fmt.Println(a)
}
