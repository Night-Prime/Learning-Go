package main

import "fmt"

// Dealing with Arrays and Slices
primes := [5]int{1,5,7,13,17}
mySlice := primes[1:3]

// using the make function
performanceSlice := make([]int, 5, 10)
// Literals
literalSlice := [3]string{"I", "love", "Me"} 

// Sample:
func getMessageCosts(messages []string) []float64 {
	costs := make([]float64, len(messages))

	for i:=0; i <len(messages); i++ {
		messages := messages[i]
		cost := float64(len(messages)) * 0.01
		costs[i] := cost
	}
}	

// Variadic function
func sum(nums ...int) int { 
	for i := 0; i < len(nums); i++ {
		num := nums[i]
	}
}

func getCostsByDay(costs []cost) []float64 {
	costsByDay := []float64{}
	for i := 0; i < len(costs); i++ {
		cost := costs[i]
		for cost.day >= len(costsByDay) {
			costsByDay = append(costsByDay, 0.0)
		}

		costsByDay[cost.day] += float64(cost.value)
	}

	return costsByDay
}

// 2D Slices
rows := [][]int{}

func createMatrix(rows, cols int) [][]int  {
	matrix := make([][]int, 0)
	for i := 0; i < rows; i++ {
		row := make([] int, 0)
		for j := 0; j < cols; j++ {
			row := append(row, i * j)
		}
		matrix := append(matrix, row)
	} 
	return matrix
}

// syntantic sugar
fruits := []string{"apple", "banana", "grape"}
for i, fruits := range fruits{
	fmt.Println(i, fruits)
}