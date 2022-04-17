package helper

import (
	"fmt"

	"github.com/aburtasov/booking-app/internal/userdata"
)

func GreetUsers(conferenceName string, conferenceTickets, remainingTickets uint) {
	fmt.Println("Welcome to our conference!")

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

}

func GetFirstNames(bookings *[]userdata.UserData) []string {

	firstNames := []string{}
	for _, booking := range *bookings {

		firstNames = append(firstNames, booking.FirstName)

	}

	return firstNames
}

func GetUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Print("Enter your first name: ")
	fmt.Scanln(&firstName)

	fmt.Print("Enter your last name: ")
	fmt.Scanln(&lastName)

	fmt.Print("Enter your email: ")
	fmt.Scanln(&email)

	fmt.Print("Enter number of tickets: ")
	fmt.Scanln(&userTickets)

	return firstName, lastName, email, userTickets
}

func BookTicket(userTickets uint, firstName string, lastName string, email string, bookings *[]userdata.UserData, remainingTickets *uint, conferenceName string) {
	*remainingTickets = *remainingTickets - userTickets

	// create a map for a user

	var userData = userdata.UserData{
		FirstName:       firstName,
		LastName:        lastName,
		Email:           email,
		NumberOfTickets: userTickets,
	}

	*bookings = append(*bookings, userData)

	fmt.Printf("List of bookings is %v\n", *bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets.You will recieve a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", *remainingTickets, conferenceName)
}
