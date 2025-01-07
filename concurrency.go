package main

import (
	"fmt"
	"time"
)

// the keyword "go" handles concurrent routine

func sendEmail(message string) {
	go func() {
		time.Sleep(time.Millisecond * 250)
		fmt.Printf("Email received:  '%s'\n", message)
	}()
	fmt.Printf("Email Sent: '%s'\n", message)
}

// using channels
func filterOldEmails(emails []email) {
	isOldChan := make(chan bool)

	go func() {
		for _,e :=range emails{
			if e.date.Before(time.Date(2020, 0,0,0,0,0,0, time.U)) {
					isOldChan <- true
					continue
			}
			isOldchan <- false 
		}
	}()

	isOld := <- isOldChan
	fmt.Println("email 1 is Old: ", isOld)
	isOld = <- isOldChan
	fmt.Println("email 2 is Old: ", isOld)
	isOld = <- isOldChan
	fmt.Println("email 3 is Old: ", isOld)
	close(ch)
}
