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

func (account *Account) LastPayment() (Payment, error) {
	if len(account.Payments) == 0 {
		return Payment{}, errors.New("Account has no payments")
	}

	lastPayment := account.Payments[0]
	for _, payment := range account.Payments {
		if payment.Date.After(lastPayment.Date) {
			lastPayment = payment
		}
	}

	return lastPayment, nil
}
