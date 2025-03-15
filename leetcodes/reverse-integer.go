package main

import (
	"fmt"
	"math"
)

// note: i can use this to solve palindrome questions

func reverseInteger(value int) int{
	// first we store the sign (making it an absolute value)
	sign := 1
	if x < 0 {
		sign = -1
		x = -x
	}

	// then we reverse the value by looping
	reversed := 0
	for x > 0 {
		digit := x % 10
		reversed = reversed * 10 + digit
		x = x / 10
	}

	// then we apply the sign back
	reversed *= sign

	// and check for overflow
	if reversed < math.MinInt32 || reversed > math.MaxInt32 {
		return 0
	}

	return reversed

}