package main

import (
	"fmt"

	"example.com/fileops"
)

const BALANCEFILE = "balance.txt"

func main(){
	displayGreeting()
	var accountBalance, err = fileops.ReadBalanceFile(BALANCEFILE)

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


func checkBalance(accountBalance float64) {
	fmt.Printf("Your balance: %.3f\n", accountBalance)
}

func depositoMoney(accountBalance *float64){
	var value float64
	fmt.Print("Input your money: ")
	fmt.Scan(&value)
	updatedAccountBalance := *accountBalance + value
	*accountBalance = updatedAccountBalance
	fileops.WriteBalanceFile(*accountBalance, BALANCEFILE)
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
	fileops.WriteBalanceFile(*accountBalance, BALANCEFILE)
}