package main

import (
	"fmt"
	"math"	
)

type Shape Interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	Height, Width float64
}

func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * (c.Radius * c.Radius)
}

func (c Circle) Perimeter() float64 {
	return  2 * math.Pi * c.Radius
}

// util function
func PrintShapeDetails(s Shape) {
	fmt.Printf("Area : %.2f \n, Perimeter : %.2f \n", s.Area(), s.Perimeter())
}

func main() {
	circle := Circle{Radius: 8.6}
	rectangle := Rectangle{Width: 7.5, Height: 3.4}

	PrintShapeDetails(circle)
	PrintShapeDetails(rectangle)

}