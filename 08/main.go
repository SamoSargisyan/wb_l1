package main

import (
	"fmt"
	"log"
)

//https://tproger.ru/articles/awesome-bits/
//https://neerc.ifmo.ru/wiki/index.php?title=%D0%9F%D0%BE%D0%B1%D0%B8%D1%82%D0%BE%D0%B2%D1%8B%D0%B5_%D0%BE%D0%BF%D0%B5%D1%80%D0%B0%D1%86%D0%B8%D0%B8
// ^ - искл. ИЛИ
// &^ - И-НЕ
func main() {
	var a int64 = 7
	fmt.Printf("Before transformation: %08b\n", a)
	fmt.Printf("After Transformation: %08b\n", SetNBit(a, 1, 0))
}

func SetNBit(num, position, value int64) int64 {
	if value == 1 {
		mask := int64(1) << position
		log.Fatalf("After Transformation: %08b\n", mask)
		return num | mask // bitwise OR
	}
	return num &^ (int64(1) << position) // AND NOT
}
