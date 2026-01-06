package main

import (
	"fmt"

	"github.com/prajna/bankapp/service"
)

func main() {
	account := service.NewAccount("Prajna")

	account.Deposit(1000)
	account.Withdraw(300)

	fmt.Println("Final Balance:", account.Balance)
}
