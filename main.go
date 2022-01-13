package main

import (
	"fmt"
	"strings"
	"time"
)

var confName = "Go Conference"

const confTickets = 50

var remainingTickets = 50
var bookings = make([]UserData, 0) // initialize a  list of structs

type UserData struct { // create a new type with specified name
	firstName       string
	lastName        string
	email           string
	numberOfTickets int
}

func main() {
	greetUsers()

	for {
		firstName, lastName, email, userTickets := getUserInput()

		isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTickets)

		// check if remainingTickets more than a user wants to buy
		if isValidName && isValidEmail && isValidTicketNumber {
			bookTicket(userTickets, firstName, lastName, email)
			// Start a new thread in the background
			go sendTicket(userTickets, firstName, lastName) // Goroutine -break out of the main thread and do this in a separate one
			printFirstNames()

			noTicketsRemaining := remainingTickets == 0
			if noTicketsRemaining {
				fmt.Println("The conference is booked out. Come back another time :)")
				break
			}
		} else {
			if !isValidName { // not valid name
				fmt.Println("First name or Last name your entered is too short.")
			}
			if !isValidEmail {
				fmt.Println("Email address you entered is not correct.")
			}
			if !isValidTicketNumber { // not valid name
				fmt.Println("Number of tickets you entered is not valid.")
			}

		}
	}

}

func greetUsers() {
	fmt.Println("Welcome to our booking app!")
	fmt.Printf("Get your tickets to attend %v \n", confName)
	fmt.Println("We have total of ", confTickets, "tickets and", remainingTickets, "are still available.")

}

func printFirstNames() { // slice of strings
	firstNames := []string{} // slice
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	fmt.Printf("THe First names of bookings are %v \n", firstNames)
}

func validateUserInput(firstName string, lastName string, email string, userTickets int) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets
	isValidEmail := strings.Contains(email, "@")

	return isValidName, isValidTicketNumber, isValidEmail
}

func getUserInput() (string, string, string, int) {
	var firstName string
	var lastName string
	var email string
	var userTickets int
	// ask user to enter data
	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName) // & stands for pointer in memory

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter Your Email: ")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets you want to buy: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets int,
	firstName string,
	lastName string,
	email string,
) {
	remainingTickets = remainingTickets - userTickets

	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	bookings = append(bookings, userData)
	fmt.Printf("List of bookings %v \n", bookings)

	fmt.Printf("%v booked %v tickets \n", firstName, userTickets)
	fmt.Printf("%v tickets remaining for %v \n", remainingTickets, confName)

}

func sendTicket(userTickets int, firstName string, lastName string) {
	time.Sleep(10 * time.Second)
	ticket := fmt.Sprintf("%v tickets for %v %v \n", userTickets, firstName, lastName) // pass string into var
	fmt.Println("--------")
	fmt.Printf("Sending ticket for:\n %v", ticket)
	fmt.Println("--------")
}
