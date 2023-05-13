package main

import (
	"fmt"
	"test1/helper"
	"time"
)

var conferenceName = "Go Conference"

const conferenceTickets int = 50

var RemainingTickets uint = 50
var bookings = make([]userData, 0)

type userData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

func main() {

	greetUsers()

	for {

		firstName, lastName, email, UserTickets := getUserInput()
		isValidName, isValidEmail, isValidTicketNumber := helper.VlaidateUseInrput(firstName, lastName, email, UserTickets, RemainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber {

			bookTickets(UserTickets, firstName, lastName, email, RemainingTickets)
			go sendTicket(UserTickets, firstName, lastName, email)

			firstNames := getFirstNames()
			fmt.Printf("The first names of bookings are: %v\n", firstNames)

			//exit application if no tickets are left
			if RemainingTickets == 0 {

				//end program
				fmt.Println("our conferance is booked out of stork. come back next year.")
				break
			}

		} else {
			if !isValidName {
				fmt.Println("first name or last name you enterd is too short")

			}
			if !isValidEmail {
				fmt.Println("email address you enterd doesn't contain @ sign")

			}
			if !isValidTicketNumber {
				fmt.Println("number of tickets you enterd is invalid")
			}

		}

	}

}
func greetUsers() {
	fmt.Printf("welcome to %v booking application", conferenceName)
	fmt.Printf(" We have total of %v tickets and %v are still available\n", conferenceTickets, RemainingTickets)
}
func getFirstNames() []string {

	//print only first names
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}

	return firstNames
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var UserTickets uint

	//asking for user input
	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email address: ")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets: ")
	fmt.Scan(&UserTickets)

	return firstName, lastName, email, UserTickets

}

func bookTickets(userTickets uint, firstName string, lastName string, email string, remainingTickets uint) {
	remainingTickets = remainingTickets - userTickets

	//create a map for user
	var userData = userData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is %v\n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets. you will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("##########")
	fmt.Printf("Sending ticket:\n %v \nto email address %v\n", ticket, email)
	fmt.Println("##########")
}
