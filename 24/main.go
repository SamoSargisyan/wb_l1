package main

import (
	"fmt"
	"math"
)

type Point struct {
	x, y float64
}

func (p Point) distance(anotherPoint *Point) float64 {
	return math.Sqrt(math.Pow(p.x-anotherPoint.x, 2) + math.Pow(p.y-anotherPoint.y, 2)) // distance
}

func newPoint(x, y float64) *Point {
	p := new(Point) //returns pointer to element (zero value)
	p.x = x
	p.y = y
	return p // return pointer to element with parameters
}

func main() {
	a := newPoint(0, 0)
	b := newPoint(4, 5)
	fmt.Println(b.distance(a))
}
