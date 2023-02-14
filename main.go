package main

import (
	"fmt"
	"strings"
)

func main() {
	conferenceName := "PHPLagos"
	const totalTicketsAvailable uint = 100
	var totalTicketsRemaining uint = 100

	fmt.Println("Welcome to", conferenceName, "You are highly invited")
	fmt.Printf("There are %v tickets reamining out of %v available tickets \n", totalTicketsRemaining, totalTicketsAvailable)

	bookings := []string{} // var bookings [totalTicketsAvailable]string

	for {
		var firstName, lastName, email string
		var userTicketPurchaseCount uint = 0

		fmt.Println("Enter your first name: ")
		fmt.Scan(&firstName)

		fmt.Println("Enter your last name: ")
		fmt.Scan(&lastName)

		fmt.Println("Enter your email address: ")
		fmt.Scan(&email)

		fmt.Println("How many tickets are you buying: ")
		fmt.Scan(&userTicketPurchaseCount)

		totalTicketsRemaining -= userTicketPurchaseCount

		nameIsValid := len(firstName) >= 2 && len(lastName) >= 2
		emailIsValid := strings.Contains(email, "@")
		ticketPurchaseIsValid := userTicketPurchaseCount > totalTicketsRemaining

		if nameIsValid && emailIsValid && ticketPurchaseIsValid {
			if userTicketPurchaseCount > totalTicketsRemaining {
				fmt.Println("Your purchase exceeds the availabe tickets")
				break
			}

			bookings = append(bookings, firstName+" "+lastName)

			fmt.Printf("Congrats!! %v %v, You have purchased %v tickets for %v. You will receive a confirmation at %v \n", firstName, lastName, userTicketPurchaseCount, conferenceName, email)
			fmt.Printf("There are %v tickets left \n", totalTicketsRemaining)

			firstNames := []string{}
			for _, booking := range bookings {
				var names = strings.Fields(booking)

				firstNames = append(firstNames, names[0])

			}

			fmt.Printf("Bookings: %v \n", firstNames)

			if totalTicketsRemaining == 0 {
				fmt.Print("Tickets sold out, till next year")
				break
			}
		} else {
			fmt.Println("Invalid inputs")
		}

	}

}
