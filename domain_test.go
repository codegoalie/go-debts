package domain

import (
	"testing"
	"time"
)

func TestAddPayment(t *testing.T) {
	account := Account{1, "Debt Account Name", []Payment{{2, 300, 3000, time.Now()}}}
	newPayment := Payment{1, 200, 2000, time.Now()}
	account.Add(newPayment)
	lastPayment := account.Payments[len(account.Payments)-1]
	if (lastPayment != newPayment) {
		t.Errorf("MakePayment adds payment. Expected: %f; Got: %f", newPayment.Amount, lastPayment.Amount)
	}
}
