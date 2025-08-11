package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const BALANCEFILE = "balance.txt"

func readBalanceFile()(float64, error){
	balanceByte, err := os.ReadFile(BALANCEFILE)

	if err != nil{
		return 1000, errors.New("failed to read balance file")
	}

	balanceText := string(balanceByte)
	balanceFloat, err := strconv.ParseFloat(balanceText, 64)
	
	if err != nil{
		return 1000, errors.New("failed to parse balance value")
	}

	return balanceFloat, nil

}

func writeBalanceFile(balance float64) error {
	balanceText := fmt.Sprint(balance)
	return os.WriteFile(BALANCEFILE, []byte(balanceText), 0644)
}

func main(){
	displayGreeting()
	var accountBalance, err = readBalanceFile()

	if err != nil{
		fmt.Println("ERROR")
		fmt.Println("An error occured")
		panic(err)
	}

	end := false
	for !end{
		displayMenu()
		userChoice := getUserChoice()
		if userChoice == 1 {
			checkBalance(accountBalance)
		} else if userChoice == 2{
			depositoMoney(&accountBalance)
		} else if userChoice == 3{
			withdrawMoney(&accountBalance)
		} else if userChoice == 4{
			end = true
		} else {
			fmt.Println("Memu is invalid!")
		}
	}

}

func displayGreeting(){
	fmt.Println("Welcome to Go Bank")
}

func displayMenu(){
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

func checkBalance(accountBalance float64) {
	fmt.Printf("Your balance: %.3f\n", accountBalance)
}

func depositoMoney(accountBalance *float64){
	var value float64
	fmt.Print("Input your money: ")
	fmt.Scan(&value)
	updatedAccountBalance := *accountBalance + value
	*accountBalance = updatedAccountBalance
	writeBalanceFile(*accountBalance)
}

func withdrawMoney(accountBalance *float64){
	checkBalance(*accountBalance)
	var value float64
	fmt.Print("Input your money: ")
	fmt.Scan(&value)
	if *accountBalance < value{
		// panic("Your account balance is insufficient")
		fmt.Println("Your account balance is insufficient!")
		return
	}
	updatedAccountBalance := *accountBalance - value
	*accountBalance = updatedAccountBalance
	writeBalanceFile(*accountBalance)
}