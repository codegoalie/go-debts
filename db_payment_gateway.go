package main

import (
	"fmt"
	"go-debts/interfaces"
)

type dbPaymentGateway struct {
	handler interfaces.DbHandler
}

func (gateway dbPaymentGateway) fetchPaymentsByAccountID(accountID int) []payment {
	row := gateway.handler.Query(fmt.Sprintf("SELECT balance, amount, paid_at FROM payments WHERE account_id = %d order by paid_at asc", accountID))
	var balance, amount float64
	var paidAt string
	var payments []payment
	for row.Next() {
		row.Scan(&balance, &amount, &paidAt)
		payments = append(payments, payment{Balance: balance, Amount: amount, PaidAt: paidAt})
	}
	return payments
}
