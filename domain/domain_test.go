package domain

import (
	"testing"
	"time"
)

func TestAddAccount(t *testing.T) {
	debitor := Debitor{1, "Chris", []Account{}}
	newAccount := Account{1, "Debt Account Name", []Payment{{2, 300, 3000, time.Now()}}}
	debitor.Add(newAccount)
	lastAccount := debitor.Accounts[len(debitor.Accounts)-1]
	if lastAccount.ID != newAccount.ID {
		t.Errorf("Add adds account. Expected: %s; Got: %s", newAccount.Name, lastAccount.Name)
	}
}

func TestAddPayment(t *testing.T) {
	account := Account{1, "Debt Account Name", []Payment{{2, 300, 3000, time.Now()}}}
	newPayment := Payment{1, 200, 2000, time.Now()}
	account.Add(newPayment)
	lastPayment := account.Payments[len(account.Payments)-1]
	if lastPayment != newPayment {
		t.Errorf("Add adds payment. Expected: %f; Got: %f", newPayment.Amount, lastPayment.Amount)
	}
}
