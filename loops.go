package main

import "fmt"

// Learning loops

for i:=0; i < 10; i++ {
	fmt.Sprintf("The Loop now is at %v", i)
}

// sample
func bulkSend(numMessages int) float64 {
	totalCost:= 0.0
	for i:=0; i < numMessages; i++ {
		totalCost += i.0 + (0.01 * flaot64(1))
	}
	return totalCost
} 

// situation where there is no condition in the loop (could cause an infinit loop without the check on line:25)
func maxMessages(thresh float64) int{
	totalCost:=0.0
	for i:=0; ; i++ {
		totalCost += 1.0 + (0.01 * flaot64(i))
		if totalCost > thresh {
			return i
		}
	}
}

// using for-loop as a while-loop implementation
plantHeight:= 0
for plantHeight <20{
	fmt.println("Still growing from the roots, it's current height: ", plantHeight)
	plantHeight++
}
fmt.Println("Plant has grown to this Height", plantHeight, "inches")

// solving leetcodes
func fizzbuzz() int {
	for i:=0;i <= 100; i++ {
		if i % 3 == 0 && i% 5 == 0 {
			fmt.Println("fizzbuzz!")
		} else if i % 3 == 0 {
			fmt.println("fizz")
		} else if i % 5 == 0 {
			fmt.println("buzz")
		} else {
			fmt.println(i)
		}
	}
}