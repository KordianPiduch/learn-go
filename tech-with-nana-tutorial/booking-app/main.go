package main

import (
	"fmt"
	"strings"
)

func main() {
	conferenceName := "Go Conference"
	const conferenceTickes = 50
	var remainingTickets uint = 50

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total %v tickets and %v are still available\n", conferenceTickes, remainingTickets)
	fmt.Println("Get you ticket here to attend")

	bookings := []string{}
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	for {
		// promt user some questions
		fmt.Println("What is you first name?")
		fmt.Scan(&firstName)
		fmt.Println("What is you last name?")
		fmt.Scan(&lastName)
		fmt.Println("What is you email?")
		fmt.Scan(&email)
		fmt.Println("How many tickets?")
		fmt.Scan(&userTickets)

		remainingTickets -= userTickets
		bookings = append(bookings, firstName+" "+lastName)

		// fmt.Printf("bookings slice full: %v\n", bookings)
		// fmt.Printf("bookings slice type: %T\n", bookings)
		// fmt.Printf("bookings slice length: %v\n", len(bookings))

		fmt.Printf("Thank you %v %v for booking %v tickets. We still have %v tickets left\n", firstName, lastName, userTickets, remainingTickets)

		onlyFirstNames := []string{}
		for _, booking := range bookings {
			var fullname = strings.Fields(booking)
			onlyFirstNames = append(onlyFirstNames, fullname[0])
		}

		fmt.Printf("First names of bookings in the application: %v\n", onlyFirstNames)

		if remainingTickets == 0 {
			fmt.Println("No more tickets. Try next year")
			break
		}
	}
}
