package fileops

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func WriteBalance(accountBalanceFile string, balance float64) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile(accountBalanceFile, []byte(balanceText), 0644)
}
func GetBalance(accountBalanceFile string) (float64, error) {
	data, err := os.ReadFile(accountBalanceFile)
	if err != nil {
		return 10000, errors.New("Failed to locate the File")
	}
	balanceText := string(data)
	balance, err := strconv.ParseFloat(balanceText, 64)
	if err != nil {
		return 10000, errors.New("Failed to aprse stored balance")
	}
	return balance, nil
}
