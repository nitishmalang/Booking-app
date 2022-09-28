package main

import (
	"fmt"
	"strings"
)

func main() {
    const conferenceTickets int = 50
    var remainingTickets int = 50
    conferenceName := "Go Conference"
	booking := []string{}

	fmt.Printf("Welcome to %v booking application . \n We have total %v tickets out of which %v tickets are still remaining", conferenceName , conferenceTickets , remainingTickets)

	for {
    var firstName string
	var lastName string
	var email string
	var userTickets int

	fmt.Println("Enter your firstName:\n")
	fmt.Scan(&firstName)
	fmt.Println("Enter your lastName:\n")
	fmt.Scan(&lastName)
	fmt.Println("Enter your email: \n")
	fmt.Scan(&email)
	fmt.Println("Enter no of tickets you wanna book\n")
	fmt.Scan(&userTickets)

	remainingTickets = remainingTickets - userTickets
	booking = append(booking, firstName + " " + lastName)

	fmt.Printf("Thank you %v %v for booking %v tickets.You will receive confirmation email at %v \n" , firstName , lastName , userTickets , email)
	fmt.Printf("%v tickets are remaining for %v \n", remainingTickets , conferenceName)

	firstName := []string{}
	for _, booking := range booking {
    var names = strings.Fields(booking)
	firstName = append(firstName , names[0])
	}
	fmt.Printf("The first names are %v \n" , firstName)


   }


   










}
  









