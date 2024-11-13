package main

import "fmt"

func main() {
	var name = "Go Conference"
	var ticketremain uint = 50
	const ticket = 50

	fmt.Printf("welcome to %v book app!\n", name)
	fmt.Printf("we have total %v tickets! %v tickets are left!\n", ticket, ticketremain)
	fmt.Println("get your tickets now!")

	var firstname string
	var lastname string
	var email string
	var userticket uint

	fmt.Println("enter your first name:")
	fmt.Scan(&firstname)
	fmt.Println("enter your last name:")
	fmt.Scan(&lastname)
	fmt.Println("enter your email:")
	fmt.Scan(&email)

	fmt.Println("number of tickets ordered:")
	fmt.Scan(&userticket)

	ticketremain = ticketremain - userticket
	fmt.Printf("Nice! user %v have ordered %v tickets! %v ticket remains %v!", firstname, userticket, name, ticketremain)
	fmt.Printf("ticket details will be sent to %v", email)

}
