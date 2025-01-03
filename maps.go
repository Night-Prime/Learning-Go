package main

import (
	"fmt",
	"errors"
)

// Learning about Map
ages := make(map[string]int)
ages["John"] = 37
ages["Mary"] = 24
ages["Martha"] = 21

// or
ages = map[string]int{
	"John": 37
	"Mary": 21
}

// sample implementation
func getUserMap(names []string, phoneNumbers []int) (map[string]user, error) {
	userMap := make(map[string]user)
	if len(names) != len(phoneNumbers) {
		return nil, errors.New("Invalid Sizes")
	}

	for i := 0; i < len(names); i++ {
		name := names[i]
		phoneNumber := phoneNumbers[i]
		userMap[name] = user {
			name: name,
			phoneNumber: phoneNumber
		}
	}
	return userMap, nil
}