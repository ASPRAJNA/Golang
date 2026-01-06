package service

import (
	"errors"

	"github.com/prajna/bankapp/utils"
)

type Account struct {
	Owner   string
	Balance float64
}

func NewAccount(owner string) *Account {
	utils.Log("Creating new account for " + owner)
	return &Account{
		Owner:   owner,
		Balance: 0,
	}
}

func (a *Account) Deposit(amount float64) error {
	if amount <= 0 {
		return errors.New("deposit amount must be positive")
	}
	a.Balance += amount
	utils.Log("Deposited amount")
	return nil
}

func (a *Account) Withdraw(amount float64) error {
	if amount > a.Balance {
		return errors.New("insufficient balance")
	}
	a.Balance -= amount
	utils.Log("Withdraw successful")
	return nil
}
