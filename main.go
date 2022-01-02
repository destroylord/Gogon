package main

import (
	"fmt"
	"gogon/helper"
	"strconv"
)

const confTicket int = 50

var name = "Dafrin"

const tiket int = 50

var remainingTickets uint = 50
var bookings = make([]map[string]string, 0)

func main() {

	// refrensi : https://pkg.go.dev/fmt#pkg-overview

	greetUsers()

	for {

		// call function getUserInput
		firstName, lastname, email, userTiket := getUserInput()
		// call function validation
		isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastname, email, userTiket, remainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber {

			// call books ticket
			bookTicket(userTiket, firstName, lastname, email)

			// call function print first name
			firstNames := getFirstName()
			fmt.Printf("There first names of bookings are: %v\n", firstNames)

			if remainingTickets == 0 {
				// end program
				fmt.Println("Our conference is booked out. Come back next year.")
				break
			}
		} else {

			if !isValidName {
				fmt.Println("first name or last name you entered is too short")
			}

			if !isValidEmail {
				fmt.Println("email addre you entered doesn't containt @ sign")
			}

			if !isValidTicketNumber {
				fmt.Printf("number of tickets you entered is invalid")
			}
		}

	}
}

func greetUsers() {
	fmt.Printf("Hello World %v\n", name)
	fmt.Printf("Tiket untuk nonton bola: %v dan tiket %v\n", confTicket, remainingTickets)
	fmt.Printf("Get your tickets here to attend")

}

func getFirstName() []string {
	firstNames := []string{}
	for _, booking := range bookings {

		firstNames = append(firstNames, booking["firstName"])
	}
	// fmt.Printf("There first names of bookings are: %v\n", firstNames)

	return firstNames
}

// validate

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastname string
	var email string
	var userTiket uint

	fmt.Printf("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Printf("Enter your last name: ")
	fmt.Scan(&lastname)

	fmt.Printf("Enter your email: ")
	fmt.Scan(&email)

	fmt.Printf("Enter number of ticket: ")
	fmt.Scan(&userTiket)

	return firstName, lastname, email, userTiket
}

func bookTicket(userTiket uint, firstName string, lastname string, email string) {
	remainingTickets = remainingTickets - uint(userTiket)

	// create a map for a user

	var userData = make(map[string]string)
	userData["firstname"] = firstName
	userData["lastname"] = lastname
	userData["email"] = email
	userData["numberOfTickets"] = strconv.FormatUint(uint64(userTiket), 10)

	bookings = append(bookings, userData)
	fmt.Printf("List of Bookings is %v\n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation eamil at %v\n",
		firstName, lastname, userTiket, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, name)
}
