package main

import (
	"fmt"
)

func removeElemOrdered(Collection []int, ind int) {
	copy(Collection[ind:], Collection[ind+1:]) // in slice since ind elem copy all elems after ind
	Collection[len(Collection)-1] = 0          // last elem is "deleted" with 0 in it
}

func removeElemNotOrdered(Collection []int, ind int) {
	Collection[ind] = Collection[len(Collection)-1] // swap deleted elem with last
	Collection[len(Collection)-1] = 0               // put zero in last elem
}

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7}
	removeElemOrdered(a, 3) // keeps order in collections but slower
	fmt.Println(a)
	a = []int{1, 2, 3, 4, 5, 6, 7}
	removeElemNotOrdered(a, 3) // doesn't keep order but faster
	fmt.Println(a)
}
