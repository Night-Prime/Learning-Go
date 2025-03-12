package main

import "fmt"

func Describe(v interface{}) {
	switch val := v.(type)
case int:
	fmt.Println("The type Ineteger: ", val)
case string:
	fmt.Println("The type String: ", val)
case bool:
	fmt.Println("The type Boolean: ", val)
case float64:
	fmt.Println("The type Floating Point: ", val)
default:
	fmt.Println("Type doesn't exist here")
}

func main() {
	intValue := Describe(23)
	stringVal := Describe("Hello People!")
	boolVal := Describe(true)
	floatVal := Describe(0.678)
}