package main

import (
	"fmt"
	"math"
)
// =====
type Shape interface {
	Area() float64
	Name() string
}
// =====
type Square struct {
	sideLength float64
}
type Rectangle struct {
	sideALength float64
	sideBLength float64
}
type Circle struct {
	radius float64
}
// =====
func (s Square) Area() float64 {
	return math.Pow(s.sideLength, 2)
}
func (s Square) Name() string {
	return "Square"
}

func (s Rectangle) Area() float64 {
	return s.sideALength * s.sideBLength
}
func (s Rectangle) Name() string {
	return "Rectangle"
}

func (s Circle) Area() float64 {
	return math.Pi * math.Pow(s.radius, 2)
}
func (s Circle) Name() string {
	return "Circle"
}

func main() {
	square := Square{10}
	rectangle := Rectangle{2, 3}
	circle := Circle{5}

	printShapeArea(square)
	printShapeArea(rectangle)
	printShapeArea(circle)
}

func printShapeArea(s Shape) {
	fmt.Printf("%s area: %.2f sm^2\n", s.Name(), s.Area())
}

