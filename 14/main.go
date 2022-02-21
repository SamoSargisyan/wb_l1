package main

import (
	"fmt"
	"reflect"
)

func main() {
	tst := "string"
	fmt.Println(reflect.TypeOf(tst))
}
