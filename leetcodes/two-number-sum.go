// package main

// import "fmt"

// func reverseNum(x []int, y []int) int {
// 	//writing a reusable anon function to reverse the slice
// 	sliceReverse := func(s []int) []int {
// 		reverse := make([]int, len(s))
// 		for i := 0; i < len(s); i++ {
// 			reverse[i] = s[len(s)-1-i]
// 		}
// 		fmt.Println("The reversed order: ", reverse)
// 		return reverse
// 	}

// 	reverseX := sliceReverse(x)
// 	reverseY := sliceReverse(y)

// 	// then convert to number:

// 	toNumber := func(s []int) int {
// 		num := 0
// 		for _, v := range s {
// 			num = num*10 + v
// 		}
// 		fmt.Println("The number here: ", num)
// 		return num
// 	}
// 	total := toNumber(reverseX) + toNumber(reverseY)
// 	return total

// }

// func main() {
// 	result := reverseNum([]int{2, 4, 3}, []int{5, 6, 4})
// 	fmt.Printf("Answer is: %d", result)
// }


// Senior Engineer approach: 

package main

import "fmt"

func addTwoNumbers(x, y []int) []int {
	fmt.Println("the values coming in: ", x, y)

	result := []int{}
	carry := 0
	i,j := 0,0

	// iterate through each slice
	for i < len(x) || j < len(y) || carry > 0 {
		sum := carry
		fmt.Printf("The sum %x at %x i:index and %x j:index \n", sum, i, j)

		if i < len(x) {
			sum += x[i]
			fmt.Printf("The sum %x for x %x  \n", sum, x)
			i++
		}

		if j < len(y) {
			sum += y[j]
			fmt.Printf("The sum %x for y %x \n", sum, y)
			j++
		}

		newDigit := sum % 10
		fmt.Printf("what's happening here: %d \n", newDigit)
		carry = sum / 10
		fmt.Printf("what's happening @carry: %d \n", carry)

		result = append(result, newDigit)
		fmt.Printf("what's happening @result: %d \n", result)
	}

	return result
}

func main() {
    x := []int{2, 4, 3} // Represents 342
    y := []int{5, 6, 4} // Represents 465
    sum := addTwoNumbers(x, y)
    fmt.Println("Answer (reversed slice):", sum) // Output: [7 0 8]
}
