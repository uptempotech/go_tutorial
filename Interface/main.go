package main

import (
	"fmt"
	"math"
)

// STRUCTS AND INTERFACES

func main() {
	rect := Rectangle{
		20,
		50,
	}
	cir := Circle{
		4,
	}

	fmt.Println("Rectangle Area =", getArea(rect))
	fmt.Println("Circle Area =", getArea(cir))
}

// An interface defines a list of methods that a type must implement.
// If that type implements those methods the proper method is executed
// even if the original is referred to with the interface name.

// Shape interface
type Shape interface {
	area() float64
}

// Rectangle defines the width and height
type Rectangle struct {
	height float64
	width  float64
}

// Circle defines the radius
type Circle struct {
	radius float64
}

func (r Rectangle) area() float64 {
	return r.height * r.width
}

func (c Circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func getArea(shape Shape) float64 {
	return shape.area()
}
