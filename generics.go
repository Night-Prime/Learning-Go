package main

import "fmt"

// Generics(Type Parameter) are simply reuseable functions in Go

func genericSample[T any](s []T) ([]T, []T) {
	mid := len(s)/2
	return s[:mid], s[:mid]
}

func main() {
	firstInts, secondInts := genericSample([]int{0,1,2,3,4,5,6})
	fmt.Println(firstInts, secondInts)
}