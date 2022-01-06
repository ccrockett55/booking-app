package main

import "fmt"

func main() {
	var confName = "Go Conference"
	const confTickets = 50
	var remainingTickets = 50

	fmt.Printf("Welcome to the %v booking application\n", confName)
	fmt.Printf("We have a total of %v tickets and there are %v remaining tickets.\n", confTickets, remainingTickets)
	fmt.Println("You can buy your tickets here.")

	var userName string
	var userTickets int

	userName = "John"
	userTickets = 2
	fmt.Printf("User %v booked %v tickets.\n", userName, userTickets)

}
