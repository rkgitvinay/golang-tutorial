package main

import (
	"fmt"
	"math"
)

// 1- Define Interface
type Shape interface {
	Area() float64
	Perimeter() float64
}

// 2.1 - Implementing an Interface Using Circle
type Circle struct {
	Radius float64
}

// Implement Area() method for Circle
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Implement Perimeter() method for Circle
func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

// 2.2 - Implementing an Interface using Ractangle
type Rectangle struct {
	Width  float64
	Height float64
}

// Implement Area() method for Rectangle
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Implement Perimeter() method for Rectangle
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Height + r.Width)
}

// Using an Interface
func PrintShapeDetails(s Shape) {
	fmt.Println("Area ", s.Area())
	fmt.Println("Perimeter ", s.Perimeter())
}

func main() {
	c := Circle{Radius: 2.5}

	r := Rectangle{Height: 2, Width: 5}

	PrintShapeDetails(c) // Circle satisfies the Shape
	PrintShapeDetails(r) // Ractangle satisfies the Shape
}
