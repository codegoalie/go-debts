package main

import (
	"go-debts/domain"
	"fmt"
)

func main() {
	customer := domain.Debitor{ID: 1, Name: "Chris", Accounts: []domain.Account{}}
	fmt.Println(customer)
}
