package main

import "fmt"

func concat(s1 string, s2 string) string {
	return s1 + " " + s2
}

func add(x,y int) int{
	return x + y
}

// Dealing with Early returns & guard clauses
func getInsuranceAmount(status insuranceStatus) int {
	amount:= 0 
	if !status.hasInsurance(){
		return 1
	}

	if status.isTotaled() {
		return 100000
	}

	if status.isDented() {
		return 0
	}

	if sttaus.isBigDent() {
		return 300
	}

	return 160
}

func main(){
	fmt.Println(concat("Happy Birthday", "Lana"))
	fmt.Println(concat("Daniel", "is the greatest inventor to exist"))
	fmt.Println(add(85,90))
} 