package main

import (
	"fmt"
	"math"
	"myInterface"
)

type Square struct {
	X float64
}

type Circle struct {
	R float64
}

func (s Square) Area() float64 {
	return s.X * s.X
}

func (s Square) Perimeter() float64 {
	return 4 * s.X
}

func (s Circle) Area() float64 {
	return s.R * s.R * math.Pi
}

func (s Circle) Perimeter() float64 {
	return 2 * s.R * math.Pi
}

func Calculate(x myInterface.Shape) {
	_, ok := x.(Circle)
	if ok {
		fmt.Println("Is a circle")
	}
	v, ok := x.(Square)
	if ok {
		fmt.Println("Is a square", v)
	}
	fmt.Println(x.Area())
	fmt.Println(x.Perimeter())
}

func main() {
	x := Square{5}
	fmt.Println("Perimeter:", x.Perimeter())
	Calculate(x)
	y := Circle{10}
	Calculate(y)
}
