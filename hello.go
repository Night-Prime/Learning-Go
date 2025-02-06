package main

import (
	"fmt"
)

func main() {
	// A FizzBuzz Program - Prints numbers from 1 to 100
	// If a number is divisible by 3, print "Fizz"
	// If a number is divisible by 5, print "Buzz"
	// If a number is divisible by both 3 and 5, print "FizzBuzz"
	// Otherwise, print the number itself

	fizz := 3  // Divisibility factor for "Fizz"
	buzz := 5  // Divisibility factor for "Buzz"

	for i := 1; i <= 100; i++ { // Loop from 1 to 100
		if i%fizz == 0 && i%buzz == 0 { // Divisible by both 3 & 5
			fmt.Println("FizzBuzz")
		} else if i%fizz == 0 { // Divisible by 3 only
			fmt.Println("Fizz")
		} else if i%buzz == 0 { // Divisible by 5 only
			fmt.Println("Buzz")
		} else { // Not divisible by 3 or 5
			fmt.Println(i)
		}
	}
}
// func main() {
// 	// Define user inputs
// 	var userName string = "Alice"
// 	message := "Hello, 世界! 🎉\nHope you're having a great day!"

// 	// Using double quotes for escaped characters (\n, \t)
// 	formattedMessage := fmt.Sprintf("User: \"%s\"\nMessage: \"%s\"", userName, message)

// 	// Using backticks for multi-line raw string
// 	chatTemplate := `------ Chat Log ------
// %s
// ----------------------`

// 	// Display the formatted chat message
// 	fmt.Printf(chatTemplate, formattedMessage) 
// 	// console: 
// 	// ------ Chat Log ------
// 	// User: "Alice"
// 	// Message: "Hello, 世界! 🎉
// 	// Hope you're having a great day!"
//     // ----------------------%  
// }

// func main() {
// 	// Boolean variables
// 	var isAdmin bool = false      // Is the user an admin?
// 	hasCorrectPasscode := true  // Did the user enter the correct passcode?
// 	isMaintenanceMode := false  // Is the system under maintenance?

// 	// Logical check: Grant access only if:
// 	// - The user is an admin OR has the correct passcode
// 	// - The system is NOT in maintenance mode
// 	canAccess := (isAdmin || hasCorrectPasscode) && !isMaintenanceMode

// 	// Output result
// 	if canAccess {
// 		fmt.Println("✅ Access Granted: Welcome to the secure vault!")
// 	} else {
// 		fmt.Println("❌ Access Denied: Please check your credentials or try later.")
// 	}
// }

// func main() {
// 	// Float:
// 	var pi float64 = 3.16
// 	thirdNum := float32(90.68)
// 	fourthNum := float64(0.000000000000078)

// 	// 1. Calculate the circumference of a circle (C = π * diameter)
// 	diameter := 10.0
// 	circumference := pi * diameter

// 	// 2. Calculate exponential growth (A = P(1 + r)^t)
// 	principal := 1000.0  // Initial amount
// 	rate := 0.05         // 5% interest rate
// 	time := 3.0          // 3 years
// 	amount := principal * math.Pow((1 + rate), time)

// 	// Print results
// 	fmt.Printf("Circumference of the circle: %.2f\n", circumference) // Circumference of the circle: 31.60
// 	fmt.Printf("Exponential growth amount: %.2f\n", amount) // Exponential growth amount: 1157.63

// 	// Print individual float values
// 	fmt.Println("Other float values:", thirdNum, fourthNum) // Other float values: 90.68 7.8e-14
// }

// func main() {
// 	// Unsigned integers:
// 	var firstNum uint = 56
// 	secondNum := uint8(245)
// 	thirdNum := uint16(64678)
// 	fourthNum := uint32(675869606)
// 	fifthNum := uint64(8596958859555869)

// 	// Convert to uint64 for safe calculations
// 	sum := uint64(firstNum) + uint64(secondNum) + uint64(thirdNum)
// 	product := uint64(thirdNum) * fifthNum
// 	difference := fifthNum - uint64(fourthNum)

// 	// Print results
// 	fmt.Printf("Sum: %d\n", sum) // Sum: 64979
// 	fmt.Printf("Product: %d\n", product) // Product: 2631782907067946702
// 	fmt.Printf("Difference: %d\n", difference) // Difference: 8596958183686263
// }


// func main() {
// 	// integers :
// 	var firstNum int = 8
// 	secondNum := int8(126)
// 	thirdNum := int16(4600)
// 	fourthNum := int32(-789586979)
// 	fifthNum := int64(67789708087860)

// 	// now imagine i intend to add some of these variables together:
// 	sumTotal := int(secondNum) + int(firstNum) + int(fourthNum)
// 	// the variables might be of different size but the int() converts
// 	// them into one data size.

// 	// Multiplication should use int64
// 	multiply := int64(thirdNum) * fifthNum

// 	fmt.Printf("The total is: %d , while the product is: %d", sumTotal, multiply)
// 	// console: The total is: -789586845 , while the product is: 311832657204156000% 
// }

// func main() {
// 	var initialVar string = "Dhaniel"
// 	const rigidVar string = "constants"

// 	initialVar = "flexible"
// 	// this is valid and would change

// 	rigidVar = "not-flexible"
// 	// this is not valid as it unable to be reassigned.

// 	fmt.Println(initialVar, rigidVar); 
// }  