package main

import "fmt"

// Dealing with Higher Order Functions - functions that takes other functions as argument
func add(x, y int) int{
	return x + y
}

func mul(x, y int) int{
	return x * y
}
// takes three inputs and a function and executes them
func  (a,b,c int, arithmetic func(int, int) int) int{
	return arithmetic(arithmetic(a, b), c)
}

func main() {
	fmt.Println(aggregate(2,3,6, add))
	fmt.Println(aggregate(2,2,4, mul))
}