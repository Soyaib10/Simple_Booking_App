package main

import (
	"fmt"
	"strconv"
)

// variables
const conferenceTickets int = 50

var conferenceName string = "GO CONFERENCE"
var remainingTickets int = 50
var bookings = make([]map[string]string, 0)

func greetUsers() {
	fmt.Printf("Welcome to %v booking application", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still availabe.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets!")
}

func getUserInput() (string, string, string, int) {
	// declare variables
	var firstName string
	var lastName string
	var email string
	var userTickets int

	// ask user their personal info
	fmt.Print("Enter your first name: ")
	fmt.Scan(&firstName)
	fmt.Print("Enter your last name: ")
	fmt.Scan(&lastName)
	fmt.Print("Enter your email: ")
	fmt.Scan(&email)
	fmt.Print("Enter number of tickets: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, element := range bookings {
		firstNames = append(firstNames, element["firstName"])
	}
	return firstNames
}

func bookTickets(userTickets int, firstName string, lastName string, email string) {
	remainingTickets -= userTickets

	// create a map
	var userData = make(map[string]string)
	userData["firstName"] = firstName
	userData["lastName"] = lastName
	userData["email"] = email
	userData["userTickets"] = strconv.FormatInt(int64(userTickets), 10)




	bookings = append(bookings, userData)
	fmt.Printf("List of bookings are: %v\n", bookings)
	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v.\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets are remaining for %v.\n", remainingTickets, conferenceName)
}

func main() {
	// Greet Users
	greetUsers()

	for {
		// call user input func
		firstName, lastName, email, userTickets := getUserInput()

		// check if user input is valid or not
		isValidName, isValidEmail, isValidUserTicket := validUserInput(firstName, lastName, email, userTickets)

		if isValidName && isValidEmail && isValidUserTicket {
			// call book ticket function
			bookTickets(userTickets, firstName, lastName, email)

			// call the first names of the users who bought tickets
			var firstNames = getFirstNames()
			fmt.Printf("The first names of the audience who booked tickets: %v.\n", firstNames)

			// check if any tickets left
			if remainingTickets == 0 {
				fmt.Println("Our conference is booked out. Please come back next year.")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("Your first or last name is way short.")
			}
			if !isValidEmail {
				fmt.Println("Your email address doesn't contain @ sign.")
			}
			if !isValidUserTicket {
				fmt.Println("Number of tickets entered is invalid.")
			}
		}
	}
}
