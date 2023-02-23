package main

import (
	"booking-app/anotherPackage"
	"fmt"
	"strings"
)

// package level variables  - available within all functions in this package
const conferenceTickes = 50

var conferenceName string = "Go Conference"
var remainingTickets uint = 50
var bookings = []string{}

func main() {

	greetUsers()

	for {
		// promt user some questions
		firstName, lastName, email, userTickets := getUserInputs()
		isValidName, isValidEmail, isValidTicketNumber := anotherPackage.ValidateInputs(firstName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber {
			bookTickets(firstName, lastName, userTickets)
			printFirstNames(bookings)

			if remainingTickets == 0 {
				fmt.Println("No more tickets. Try next year")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("First Name or Last Name is too short")
			}
			if !isValidEmail {
				fmt.Println("Email address does not contain @ symbol")
			}
			if !isValidTicketNumber {
				fmt.Println("Number of tickets is invalid")
			}
			continue
		}

	}
}

func getUserInputs() {
	panic("unimplemented")
}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total %v tickets and %v are still available\n", conferenceTickes, remainingTickets)
	fmt.Println("Get you ticket here to attend")
}

func printFirstNames(bookings []string) {
	onlyFirstNames := []string{}
	for _, booking := range bookings {
		var fullname = strings.Fields(booking) // works like split string in python
		onlyFirstNames = append(onlyFirstNames, fullname[0])
	}
	fmt.Printf("First names of bookings in the application: %v\n", onlyFirstNames)
}

func bookTickets(firstName string, lastName string, userTickets uint) {
	remainingTickets -= userTickets
	bookings = append(bookings, firstName+" "+lastName)

	// fmt.Printf("bookings slice full: %v\n", bookings)
	// fmt.Printf("bookings slice type: %T\n", bookings)
	// fmt.Printf("bookings slice length: %v\n", len(bookings))

	fmt.Printf("Thank you for booking %v tickets. We still have %v tickets left\n", userTickets, remainingTickets)
}
