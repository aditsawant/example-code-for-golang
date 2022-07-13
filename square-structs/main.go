package main

import (
	"fmt"
)

func main() {
	fmt.Println("Working with Square structs")

	sq, err := NewSquare(0, 0, 4)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", sq)
	fmt.Println("The area of the square: ", sq.GetArea())

	sq.Move(2, 2)

	fmt.Printf("%+v\n", sq)
	fmt.Println("The area of the square: ", sq.GetArea())
}

func NewSquare(x, y, len int) (*Square, error) {
	if len <= 0 {
		err := fmt.Errorf("can't have edges with non-positive length, please try again")
		return nil, err
	}
	sq := Square{x, y, len}
	return &sq, nil
}

type Square struct {
	x   int
	y   int
	len int
}

func (sq *Square) Move(dx, dy int) {
	sq.x += dx
	sq.y += dy
}

func (sq *Square) GetArea() int {
	area := sq.len * sq.len
	return area
}
