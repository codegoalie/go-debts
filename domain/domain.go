package domain

import (
	"time"
)

type Debitor struct {
	ID       int
	Name     string
	Accounts []Account
}

type Account struct {
	ID       int
	Name     string
	Payments []Payment
}

type Payment struct {
	ID      int
	Amount  float64
	Balance float64
	Date    time.Time
}

func (debitor *Debitor) Add(account Account) error {
	debitor.Accounts = append(debitor.Accounts, account)
	return nil
}

func (account *Account) Add(payment Payment) error {
	account.Payments = append(account.Payments, payment)
	return nil
}
