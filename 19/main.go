package main

import (
	"fmt"
)

func reverseWord(inpStr string) string {
	out := make([]rune, len(inpStr)) // To avoid memory relocation
	// using runes bc of unicode is in task
	bytes := []rune(inpStr)

	counter := 0
	for i := len(bytes) - 1; i >= 0; i-- {
		out[counter] = bytes[i]
		counter++
	}
	return string(out)
}

func main() {
	test := "абвGD"
	out := reverseWord(test)
	fmt.Println(out)
}
