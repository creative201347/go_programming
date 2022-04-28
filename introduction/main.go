package main

import (
	"fmt"
	"strings"
)

var conferenceTickets uint = 50
var conferenceName = "Go Conference"
var remaningTickets uint = 50
var bookings []string

func main() {
	greetUsers()
	for {
		firstName, lastName, email, userTickets := getUserInput()
		isValidName, isValidEmail, isValidTicketNumber := ValidateUserInput(firstName, lastName, email, userTickets, remaningTickets)

		if isValidName && isValidEmail && isValidTicketNumber {
			bookTicket(userTickets, firstName, lastName, email)

			fmt.Printf("These first name of bookings are : %v\n", getFirstNames())

			if remaningTickets == 0 {
				fmt.Println("Our conference is booked out. Come back next year.")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("First name or last name is too short")
			}
			if !isValidEmail {
				fmt.Println("Not a valid email address")
			}
			if !isValidTicketNumber {
				fmt.Println("Number of tickets you've entered is invalid")
			}
		}
	}
}
func greetUsers() {
	fmt.Printf("Welcome to %v booking applications\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still avaliable\n", conferenceTickets, remaningTickets)
	fmt.Println("Get your ticket here to attend")
}
func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}
	return firstNames
}
func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email address: ")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remaningTickets = remaningTickets - userTickets
	bookings = append(bookings, firstName+" "+lastName)

	fmt.Printf("Thank you %v %v for booking %v tickets.\nYou will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remaningTickets, conferenceName)

}
