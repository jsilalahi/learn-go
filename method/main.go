package main

import (
	"fmt"
)

// Point a thing
type Point struct {
	X, Y float64
}

func (p Point) Area() float64 {
	return p.X * p.Y
}

func main() {
	p := Point{4, 2}
	fmt.Println(fmt.Sprintf("Area %f and %f = %f", p.X, p.Y, p.Area()))
}
