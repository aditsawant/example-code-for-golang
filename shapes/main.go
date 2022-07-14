package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Types implement the Shape interface")

	shapes := make([]Shape, 3)
	shapes[0] = Circle{2}
	shapes[1] = Circle{3}
	shapes[2] = Rect{2, 4}

	for _, shape := range shapes {
		fmt.Println("The area of the shape is:", shape.getArea())
	}
}

type Rect struct {
	height float64
	width  float64
}

type Circle struct {
	rad float64
}

func (r Rect) getArea() float64 {
	return r.height * r.width
}

func (c Circle) getArea() float64 {
	return c.rad * c.rad * math.Pi
}

type Shape interface {
	getArea() float64
}
