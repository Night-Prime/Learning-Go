package main

import "fmt"

// Error Interfaces
type userError struct {
	name string
}

func (e userError) Error() string{
	return fmt.Sprintf("%v has a problem with their account", e)
}

func sendSMS(msg, userName string) error {
	if !canSendToUser(username) {
		return userError{name: userName}
	}

	return  fmt.Sprintf("sending you lots of love from Dhaniel Corp", userName)
}