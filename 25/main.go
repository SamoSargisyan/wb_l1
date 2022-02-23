package main

import (
	"fmt"
	"time"
)

func handmadeSleep(s int) {
	<-time.After(time.Duration(s) * time.Second)
}

func main() {
	fmt.Println("Before sleep...")
	handmadeSleep(2)
	fmt.Println("After sleep...")
}
