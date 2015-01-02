package main

import (
	"go-debts/domain"
	"fmt"
)

func main() {
	customer := domain.Debitor{1, "Chris", []domain.Account{}}
	fmt.Println(customer)
}
