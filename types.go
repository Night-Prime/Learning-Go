package main

import "fmt"

// := replaces var and infers type

func main() {

	var smsSendingLimit int
	var costPerSMS float64
	var hasPermission bool
	var username string

	fmt.Printf(
		"%v %f %v %q\n",
		smsSendingLimit,
		costPerSMS,
		hasPermission,
		username,
	)

	congats:= "happy birthday!"
	fmt.Println(congats)
 
	penniesPerText := 2.0
	fmt.Printf("The type of penniesPerText is %T\n", penniesPerText)

	avgOpenRate, displayMessageRate := 4, "is the avaerage open rate of your message"
	fmt.Println(avgOpenRate, displayMessageRate)

	accountAge := 2.6
	accountAgeInt := int(accountAge)
	fmt.Println("Your account has existed for :", accountAgeInt, "years")

	const name = "Saul Goodman"
	const openRate = 30.5

	msg := fmt.Sprintf(
		"Hi %s your open rate is %.1f percent",
		name,
		openRate,
	)

	fmt.Println(msg)

} 