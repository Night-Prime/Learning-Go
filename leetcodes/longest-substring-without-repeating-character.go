package main

import (
	"fmt"
)

// This doesn't relate to the answer, just thought it'd be fun to try
func lengthOfSubString(word string){
	count := 0
	for i := 0; i < len(word); i++ {
		value := word[i]
		fmt.Printf("%c \n", value)
		fmt.Println("------------------------ \n")

		if(i == i+1) {
			count++
		}
		
	}
}

func lengthOfSubstringWithoutRepeatingCharacter (s string) int{
	// using the sliding window technique (deepseek)
	lastSeen := make(map[byte]int)
	left := 0
	maxLen := 0

	for right:= 0; right < len(s); right++ {
		char := s[right]
		fmt.Printf("First character index: %c \n", char);
		// when the charcter already exists in the map
		if idx, exists := lastSeen[char]; exists && idx >= left {
			left = idx + 1 // move pointer to exclude repeated character
			fmt.Printf("Next Pointer Index: @%v \n", left);
		}
		lastSeen[char] = right
		fmt.Printf("Last seen character index: @%v \n", right);
		if windowLen := right - left + 1; windowLen > maxLen {
			maxLen = windowLen
		}
	}
	fmt.Printf("Each Length: %x", maxLen);
	return maxLen
}

func main() {
	values := "word"
	lengthOfSubString(values)

	// leetcode 

	word := "aaaafghbbbbfjsi"
	lengthOfSubstringWithoutRepeatingCharacter(word)
}