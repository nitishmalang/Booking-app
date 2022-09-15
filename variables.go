package main
import "fmt"
func main() {
  const conferenceTickets int = 50
  var remainingTickets int = 50
  conferenceName := "Go Conference"

  fmt.Printf("Welcome %v booking application.\n We have total of %v tickets and %v are still available ", conferenceName , conferenceTickets , remainingTickets)
   var firstName string
   var lastName string
   var email string
   var userTickets int

   fmt.Println("Enter your first name")
   fmt.Scan(&firstName)
   fmt.Println("Enter your last name")
   fmt.Scan(&lastName)
   fmt.Println("Enter your email ")
   fmt.Scan(&email)
   fmt.Println("Enter userTickets")
   fmt.Scan(&userTickets)

   remainingTickets = remainingTickets - userTickets

   fmt.Printf("Thank you %v %v for booking %v tickets \n", firstName , lastName , userTickets)
   fmt.Printf("%v tickets are remaining for %v \n", remainingTickets , conferenceName)





}