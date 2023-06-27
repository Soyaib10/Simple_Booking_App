package helper

import "strings"

func ValidUserInput(firstName string, lastName string, email string, userTickets int, remainingTickets int) (bool, bool, bool) {
	var isValidName = len(firstName) >= 2 && len(lastName) >= 2
	var isValidEmail = strings.Contains(email, "@")
	var isValidUserTicket = userTickets > 0 && userTickets <= remainingTickets
	return isValidName, isValidEmail, isValidUserTicket
}