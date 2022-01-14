package main

import "fmt"

func main() {

	var conferenceName string = "GO CONFERENCE"
	const conferenceTicket = 50
	var remainingTicket int = 50

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still remaining\n", conferenceName, remainingTicket)
	fmt.Println("Get your ticket here to attend")

	// var userName string
	var userTicket int

	var firstName string
	var lastName string
	var email string
	var booking []string

	fmt.Printf("Enter your firstName ")
	fmt.Scan(&firstName)

	fmt.Printf("Enter your lastName ")
	fmt.Scan(&lastName)

	fmt.Printf("Enter your email ")
	fmt.Scan(&email)

	fmt.Printf("Enter number of tickets: ")
	fmt.Scan(&userTicket)

	remainingTicket = remainingTicket - 1
	booking = append(booking, firstName+" "+lastName)

	booking[0] = firstName + " " + lastName
	// fmt.Printf("The whole array: %v\n", booking)
	// fmt.Printf("The first value is %v\n", booking[0])
	// fmt.Printf("The length of the array is %v\n", len(booking))

	fmt.Printf("Thank you %v %v for booking %v tickests. You will recieve confirmation email at %v\n", firstName, lastName, userTicket, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTicket, conferenceName)
	fmt.Printf("These are all our bookings: %v", booking)

}
