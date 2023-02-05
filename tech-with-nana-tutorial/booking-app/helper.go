package main

import "fmt"

func getUserInputs() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Println("What is you first name?")
	fmt.Scan(&firstName)
	fmt.Println("What is you last name?")
	fmt.Scan(&lastName)
	fmt.Println("What is you email?")
	fmt.Scan(&email)
	fmt.Println("How many tickets?")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}
