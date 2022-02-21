package main

import (
	"fmt"
)

func arrayUnique(incomeSlice []string) []string {
	keys := make(map[string]bool)

	uniqueSlice := make([]string, 0)

	for _, entry := range incomeSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			uniqueSlice = append(uniqueSlice, entry)
		}
	}

	return uniqueSlice
}

func main() {
	x := []string{"cat", "cat", "dog", "cat", "tree"}
	y := arrayUnique(x)

	fmt.Println(y)
}
