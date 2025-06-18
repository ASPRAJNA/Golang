package main

import (
	"asprajna/bank/fileops"
	"fmt"
)

const accountBalanceFile = "balance.txt"

func main() {
	fmt.Println("Welcome to Go Bank!\n")
	var accountBalance, err = fileops.GetBalance(accountBalanceFile)
	if err != nil {
		fmt.Println("Error occured")
		fmt.Println(err)
		fmt.Println("----------------------")
		//panic("Can't continue") manual error
	}
	for i := 0; i < 200; i++ {
		presentOptions()
		var choice int
		fmt.Println("Enter choice : ")
		fmt.Scan(&choice)
		fmt.Println("You have entered the choice : ", choice)

		//wantsCheckBalance := choice ==1
		switch choice {
		case 1:
			fmt.Println("Your Balance is : ", accountBalance)
		case 2:
			fmt.Println("Your deposit ammount : ")
			var depositAmount float64
			fmt.Scan(&depositAmount)
			if depositAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0.")
				continue
			}
			accountBalance += depositAmount
			fileops.WriteBalance(accountBalanceFile, accountBalance)
		case 3:
			fmt.Println("Withdraw amount : ")
			var withdrawAmount float64
			fmt.Scan(&withdrawAmount)

			if withdrawAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0.")
				continue
			}
			if withdrawAmount >= accountBalance {
				fmt.Println("Insuffient Balance")
				fmt.Println("Check your balance once")
				continue

			} else if accountBalance-withdrawAmount < 1000 {
				fmt.Println("Minimum 1000 balance is needed ")
				fmt.Println("Check your balance once")
				continue

			} else {
				accountBalance -= withdrawAmount
				fileops.WriteBalance(accountBalanceFile, accountBalance)
				fmt.Println("Amount withdrwan %.2f and balance is %.2f", withdrawAmount, accountBalance)
			}
		default:
			fmt.Println("Goodbye!")
			fmt.Println("Thanks for choosing our bank")
			return
		}

	}
}
