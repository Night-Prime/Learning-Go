package main

import "fmt"

// Interfaces are simply collections of  method signatures
type shape interface{
	area() float64
	perimeter() float64
}

type rect struct{
	width, height float64
}

func(r rect) area() float64 {
	return r.width * r.height 
}

func(r rect) perimeter() float64 {
	return 2*r.width + 2*r.height
}