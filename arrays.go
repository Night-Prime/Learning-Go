package main

import (
	"fmt"
	"regexp"
	"unicode"
)

func checkPasswordStrength(password string) bool {
		// Check if the password is at least 8 characters long
		if len(password) < 8 {
			return false
		}
		// Check if the password contains at least one lowercase letter
		hasLower := false
		// Check if the password contains at least one uppercase letter
		hasUpper := false

		for _, char := range password {
			switch {
			case unicode.IsLower(char):
				hasLower = true
			case unicode.IsUpper(char):
				hasUpper = true
		}

		// Password must meet all criteria
		return hasLower && hasUpper
	}
	func main() {
		// Test the password strength checker
		var password string
		fmt.Print("Enter a password: ")
		fmt.Scanln(&password)

		// Check if the password is strong
		if checkPasswordStrength(password) {
			fmt.Println("Password is strong!")
		} else {
			fmt.Println("Password does not meet the required security criteria.")
		}
	}
}


// func main() {
// 	// Declare a Celsius temperature as a float
// 	var celsius float64
	
// 	// Ask the user to input a Celsius temperature
// 	fmt.Print("Enter temperature in Celsius: ")
// 	fmt.Scanln(&celsius)
	
// 	// Perform the conversion from Celsius to Fahrenheit
// 	fahrenheit := (celsius * 9 / 5) + 32
	
// 	// Print the result
// 	fmt.Printf("%.2f Celsius is %.2f Fahrenheit\n", celsius, fahrenheit)
// }

// func main() {
// 	// Method 1: Using make
//     studentScores := make(map[string]int)
    
//     // Method 2: Map literal
//     prices := map[string]float64{
//         "apple":  0.99,
//         "banana": 0.59,
//         "orange": 0.89,
//     }
// }

// type Shape interface {
//    Area() float64
//    Perimeter() float64
// }

// // Rectangle implements Shape
// type Rectangle struct {
//    Width  float64
//    Height float64
// }

// // Circle implements Shape
// type Circle struct {
//    Radius float64
// }

// // Rectangle method
// func (r Rectangle) Area() float64 {
//    return r.Width * r.Height
// }

// func (r Rectangle) Perimeter() float64 {
//    return 2 * (r.Width + r.Height)
// }

// // Circle method
// func (c Circle) Area() float64 {
//    return math.Pi * c.Radius * c.Radius
// }

// func (c Circle) Perimeter() float64 {
//    return 2 * math.Pi * c.Radius
// }


// // User represents a customer in our e-commerce platform
// type User struct {
//    ID        int
//    FirstName string
//    LastName  string
//    Email     string
//    CreatedAt time.Time
//    Orders    []Order
// }

// // Order represents a purchase made by a user
// type Order struct {
//    ID          int
//    UserID      int
//    Items       []Item
//    TotalAmount float64
//    Status      string
//    OrderDate   time.Time
// }

// // Item represents a product in an order
// type Item struct {
//    ID       int
//    Name     string
//    Price    float64
//    Quantity int
// }
// func main() {
//    // Create a new user
//    user := User{
//        ID:        1,
//        FirstName: "John",
//        LastName:  "Doe",
//        Email:     "john.doe@example.com",
//        CreatedAt: time.Now(),
//    }

//    // Create some items
//    items := []Item{
//        {
//            ID:       1,
//            Name:     "Mechanical Keyboard",
//            Price:    149.99,
//            Quantity: 1,
//        },
//        {
//            ID:       2,
//            Name:     "Gaming Mouse",
//            Price:    79.99,
//            Quantity: 2,
//        },
//    }
//     // Create a new order
//    order := Order{
//        ID:        1,
//        UserID:    user.ID,
//        Items:     items,
//        Status:    "Pending",
//        OrderDate: time.Now(),
//    }

// }

// func main() {
//     // Method 1: Using slice literal
//     numbers := []int{1, 2, 3, 4, 5}
//     fmt.Printf("Slice literal: %v\n", numbers)

//     // Method 2: Using make
//     scores := make([]int, 3, 5)  // length 3, capacity 5
//     fmt.Printf("Made slice: %v (length: %d, capacity: %d)\n", 
//         scores, len(scores), cap(scores))

//     // Method 3: Slicing an array
//     arr := [5]int{10, 20, 30, 40, 50}
//     slice1 := arr[1:4]  // elements from index 1 to 3
//     fmt.Printf("Slice from array: %v\n", slice1)

//     // Basic Operations

//     // Appending elements
//     numbers = append(numbers, 6)
//     numbers = append(numbers, 7, 8, 9)
//     fmt.Printf("After append: %v\n", numbers)

//     // Copying slices
//     src := []int{1, 2, 3}
//     dst := make([]int, len(src))
//     copy(dst, src)
//     fmt.Printf("Copied slice: %v\n", dst)

//     // 3. Slice Manipulation
//     fruits := []string{"apple", "banana", "orange", "mango", "grape"}
    
//     // Slicing
//     someFruits := fruits[1:4]
//     fmt.Printf("Sliced fruits[1:4]: %v\n", someFruits)

//     // Re-slicing
//     moreFruits := someFruits[:4]
//     fmt.Printf("Re-sliced fruits: %v\n", moreFruits)

//     // 4. Dynamic Growth
//     fmt.Println("\nDynamic Growth:")
    
//     s := make([]int, 0)
//     fmt.Printf("Initial: len=%d cap=%d\n", len(s), cap(s))
    
//     // Append elements and watch capacity grow
//     for i := 0; i < 5; i++ {
//         s = append(s, i)
//         fmt.Printf("Append %d: len=%d cap=%d\n", i, len(s), cap(s))
//     }

//     // 5. Common Slice Operations
//     fmt.Println("\nCommon Operations:")
    
//     // Remove element at index 2
//     numbers = append(numbers[:2], numbers[3:]...)
//     fmt.Printf("After removing element at index 2: %v\n", numbers)

//     // Insert element at index 1
//     numbers = append(numbers[:1], append([]int{100}, numbers[1:]...)...)
//     fmt.Printf("After inserting 100 at index 1: %v\n", numbers)

//     // Clear slice
//     numbers = numbers[:0]
//     fmt.Printf("After clearing: %v\n", numbers)
// }


// // Dealing with Arrays and Slices
// primes := [5]int{1,5,7,13,17}
// mySlice := primes[1:3]

// // using the make function
// performanceSlice := make([]int, 5, 10)
// // Literals
// literalSlice := [3]string{"I", "love", "Me"} 

// // Sample:
// func getMessageCosts(messages []string) []float64 {
// 	costs := make([]float64, len(messages))

// 	for i:=0; i <len(messages); i++ {
// 		messages := messages[i]
// 		cost := float64(len(messages)) * 0.01
// 		costs[i] := cost
// 	}
// }	

// // Variadic function
// func sum(nums ...int) int { 
// 	for i := 0; i < len(nums); i++ {
// 		num := nums[i]
// 	}
// }

// func getCostsByDay(costs []cost) []float64 {
// 	costsByDay := []float64{}
// 	for i := 0; i < len(costs); i++ {
// 		cost := costs[i]
// 		for cost.day >= len(costsByDay) {
// 			costsByDay = append(costsByDay, 0.0)
// 		}

// 		costsByDay[cost.day] += float64(cost.value)
// 	}

// 	return costsByDay
// }

// // 2D Slices
// rows := [][]int{}

// func createMatrix(rows, cols int) [][]int  {
// 	matrix := make([][]int, 0)
// 	for i := 0; i < rows; i++ {
// 		row := make([] int, 0)
// 		for j := 0; j < cols; j++ {
// 			row := append(row, i * j)
// 		}
// 		matrix := append(matrix, row)
// 	} 
// 	return matrix
// }

// // syntantic sugar
// fruits := []string{"apple", "banana", "grape"}
// for i, fruits := range fruits{
// 	fmt.Println(i, fruits)
// }