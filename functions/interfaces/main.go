package main

import (
	"fmt"
	"math"
)

type Square struct {
	lenght float64
	width  float64
}

type Circle struct {
	radius float64
}

func (s Square) area() float64 {
	area := s.lenght * s.width
	return area
}

func (c Circle) area() float64 {
	area := math.Pi * math.Pow(c.radius, 2)
	return area
}

// this means "every type that has method area is shape"
type shape interface {
	area() float64
}

func info(s shape) float64 {
	return s.area()
}

func main() {
	var a float64

	sq := Square{
		lenght: 2,
		width:  3,
	}

	a = info(sq)
	fmt.Println(a)

	cir := Circle{
		radius: 5,
	}

	a = info(cir)
	fmt.Println(a)
}
