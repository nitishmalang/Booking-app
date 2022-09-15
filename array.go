package main
import "fmt"

func main() {
  const conferenceTickets int = 50
  var remainingTickets int = 50
  conferenceName   := " Go Conference"
  bookings := []string{}
  
  fmt.Printf("Welcome %v booking application\n . We have total %v tickets and %v tickets are still available", conferenceName , conferenceTickets , remainingTickets )
  
  var firstName string
  var lastName string
  var email string
  var userTickets int

  fmt.Println("Enter your first name \n")
  fmt.Scan(&firstName)
  fmt.Println("Enter your last name \n")
  fmt.Scan(&lastName)
  fmt.Println("Enter your email \n")
  fmt.Scan(&email)
  fmt.Println("Enter no of tickets you wanna book \n")
  fmt.Scan(&userTickets)

  remainingTickets = remainingTickets - userTickets
  bookings = append(bookings, firstName + " " + lastName)
  fmt.Printf("Thank you %v %v for booking %v tickets . You will confirmation on your email %v soon \n", firstName , lastName , userTickets , email)
  fmt.Printf("%v tickets are remaining for %v\n" , remainingTickets , conferenceName)
  fmt.Printf("These are %v all our bookings" , bookings)
  









}