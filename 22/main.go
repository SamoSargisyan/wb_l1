package main

import (
	"fmt"
	"math/big"
)

func main() {
	x := big.NewInt(2097152)
	y := big.NewInt(2097152)

	fmt.Println(new(big.Int).Add(x, y))
	fmt.Println(new(big.Int).Sub(x, y))
	fmt.Println(new(big.Int).Mul(x, y))
	fmt.Println(new(big.Int).Div(x, y))
}
