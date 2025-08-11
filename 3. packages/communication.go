package main

import (
	"fmt"

	"github.com/Pallinder/go-randomdata"
)

func displayGreeting() {
	fmt.Println("Welcome to Go Bank")
	fmt.Println("Call us 24/7", randomdata.PhoneNumber())
}

func displayMenu() {
	fmt.Println("\nPlease select a menu: ")
	fmt.Println("1. Check Balance")
	fmt.Println("2. Deposito Money")
	fmt.Println("3. Withdraw Money")
	fmt.Println("4. Exit")
}

func getUserChoice() float64 {
	var userChoice float64
	fmt.Print("Your choice: ")
	fmt.Scan(&userChoice)
	return userChoice
}