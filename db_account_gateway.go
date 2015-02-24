package main

import (
	"fmt"
	"go-debts/interfaces"
)

type dbAccountGateway struct {
	handler interfaces.DbHandler
}

func (gateway dbAccountGateway) fetchAccountsByDebitorId(debitorId int) []account {
	row := gateway.handler.Query(fmt.Sprintf("SELECT id, name FROM accounts WHERE debitor_id = %d", debitorId))
	var accountId int
	var name string
	var balance float64
	var accounts []account
	for row.Next() {
		row.Scan(&accountId, &name)
		paymentRow := gateway.handler.Query(fmt.Sprintf("SELECT balance FROM payments WHERE account_id = %d ORDER BY paid_at DESC LIMIT 1", accountId))
		paymentRow.Next()
		paymentRow.Scan(&balance)
		fmt.Println("Balance: ", balance)
		accounts = append(accounts, account{ID: accountId, Name: name, Balance: balance})
	}
	return accounts
}
