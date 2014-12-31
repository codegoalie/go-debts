package domain

import (
	"time"
)

type Debitor struct {
	Id int
	Name string
	Accounts []Account
}

type Account struct {
	Id int
	Name string
	Payments []Payment
}

type Payment struct {
	Id int
	Amount float64
	Balance float64
	Date time.Time
}

func (account *Account) Add(payment Payment) error {
	account.Payments = append(account.Payments, payment)
	return nil
}
