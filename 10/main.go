package main

import (
	"fmt"
)

// https://www.codegrepper.com/code-examples/go/golang+round+float
// прикол конечно жесткий с этими "поделить на 10 и сразу умножить на 10". Ору
func tempGroups(temps []float64) map[int][]float64 {
	tempGroups := make(map[int][]float64)
	for _, v := range temps {
		tempGroups[int(v)/10*10] = append(tempGroups[int(v)/10*10], v)
	}
	return tempGroups
}

func main() {
	temps := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	out := tempGroups(temps)
	fmt.Println(out)
}
