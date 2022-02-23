package main

import (
	"fmt"
)

func isUnique(stringForCheck string) bool {
	m := make(map[rune]bool)
	for _, i := range stringForCheck {
		_, ok := m[i]
		if ok {
			return false
		}
		m[i] = true
	}

	return true
}

func main() {
	test := []string{"abcd", "abCdefAaf", "aabcd", ""}
	for _, val := range test {
		fmt.Println(isUnique(val))
	}
}
