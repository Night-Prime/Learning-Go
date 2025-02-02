package main

import "fmt" 

func main() {
	var initialVar string = "Dhaniel"
	const rigidVar string = "constants"

	initialVar = "flexible"
	// this is valid and would change

	rigidVar = "not-flexible"
	// this is not valid as it unable to be reassigned.

	fmt.Println(initialVar, rigidVar); 
}  