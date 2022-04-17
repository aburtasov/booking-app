package main

import (
	"fmt"
	"time"

	"github.com/aburtasov/booking-app/internal/helper"
	"github.com/aburtasov/booking-app/internal/userdata"
	"github.com/aburtasov/booking-app/internal/validate"
)

func main() {

	const conferenceTickets uint = 50

	var conferenceName = "Go Conference"
	var remainingTickets uint = 50
	var bookings = make([]userdata.UserData, 0)

	helper.GreetUsers(conferenceName, conferenceTickets, remainingTickets)

	for {

		firstName, lastName, email, userTickets := helper.GetUserInput()

		isValidName, isValidEmail, isValidTicketNumber := validate.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber {
			helper.BookTicket(userTickets, firstName, lastName, email, &bookings, &remainingTickets, conferenceName)
			go sendTicket(userTickets, firstName, lastName, email)

			firstNames := helper.GetFirstNames(&bookings)
			fmt.Printf("These are all our bookings: %v\n", firstNames)

			if remainingTickets == 0 {
				fmt.Println("Tickets end")
				break
			} else {
				continue
			}
		} else {
			if !isValidName {
				fmt.Println("fist name or last name you entered is too short")
			}
			if !isValidEmail {
				fmt.Println("incorrect email")
			}
			if !isValidTicketNumber {
				fmt.Println("number of tickets you entered is invalid ")
			}

		}

	}

}

func sendTicket(userTickets uint, fistName string, lastName string, email string) {
	time.Sleep(10 * time.Second)

	ticket := fmt.Sprintf("%v tickets for %v %v", userTickets, fistName, lastName)
	fmt.Println("##############")
	fmt.Printf("Sending ticket:\n %v \nto email address %v\n", ticket, email)
	fmt.Println("##############")
}
