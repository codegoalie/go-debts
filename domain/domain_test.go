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

func TestCurrentBalanceNoPayments(t *testing.T) {
	account := Account{1, "Debt Account Name", []Payment{}}
	currentBalance := account.CurrentBalance()
	if currentBalance != 0 {
		t.Errorf("CurrentBalance with no payments. Expected: %d; Got: %f", 0, currentBalance)
	}
}

func TestCurrentBalanceOnePayment(t *testing.T) {
	payment := Payment{2, 300, 3000, time.Now()}
	account := Account{1, "Debt Account Name", []Payment{payment}}
	currentBalance := account.CurrentBalance()
	if currentBalance != payment.Balance {
		t.Errorf("CurrentBalance with one payments. Expected: %f; Got: %f", payment.Balance, currentBalance)
	}
}

func TestCurrentBalanceTwoOrderedPayment(t *testing.T) {
	minusOneDay, _ := time.ParseDuration("-24h")
	firstPayment := Payment{1, 300, 3000, time.Now().Add(minusOneDay)}
	secondPayment := Payment{2, 320, 2700, time.Now()}
	account := Account{1, "Debt Account Name", []Payment{firstPayment, secondPayment}}
	currentBalance := account.CurrentBalance()
	if currentBalance != secondPayment.Balance {
		t.Errorf("CurrentBalance with two ordered payments. Expected: %f; Got: %f", secondPayment.Balance, currentBalance)
	}
}

func TestCurrentBalanceTwoUnorderedPayment(t *testing.T) {
	minusOneDay, _ := time.ParseDuration("-24h")
	firstPayment := Payment{1, 320, 2700, time.Now()}
	secondPayment := Payment{2, 300, 3000, time.Now().Add(minusOneDay)}
	account := Account{1, "Debt Account Name", []Payment{firstPayment, secondPayment}}
	currentBalance := account.CurrentBalance()
	if currentBalance != firstPayment.Balance {
		t.Errorf("CurrentBalance with two unordered payments. Expected: %f; Got: %f", firstPayment.Balance, currentBalance)
	}
}

func TestLastPaymentNoPayments(t *testing.T) {
	account := Account{1, "Debt Account Name", []Payment{}}
	_, err := account.LastPayment()
	if err == nil {
		t.Errorf("LastPayment with no payments. Expected error, but got none.")
	}
}

func TestLastPaymentOnePayment(t *testing.T) {
	payment := Payment{2, 300, 3000, time.Now()}
	account := Account{1, "Debt Account Name", []Payment{payment}}
	lastPayment, _ := account.LastPayment()
	if lastPayment != payment {
		t.Errorf("LastPayment with one payment. Expected: %f; Got: %f", payment.Amount, lastPayment.Amount)
	}
}

func TestLastPaymentTwoOrderedPayment(t *testing.T) {
	minusOneDay, _ := time.ParseDuration("-24h")
	firstPayment := Payment{1, 300, 3000, time.Now().Add(minusOneDay)}
	secondPayment := Payment{2, 320, 2700, time.Now()}
	account := Account{1, "Debt Account Name", []Payment{firstPayment, secondPayment}}
	lastPayment, _ := account.LastPayment()
	if lastPayment != secondPayment {
		t.Errorf("LastPayment with two ordered payments. Expected: %f; Got: %f", secondPayment.Amount, lastPayment.Amount)
	}
}

func TestLastPaymentTwoUnorderedPayment(t *testing.T) {
	minusOneDay, _ := time.ParseDuration("-24h")
	firstPayment := Payment{1, 320, 2700, time.Now()}
	secondPayment := Payment{2, 300, 3000, time.Now().Add(minusOneDay)}
	account := Account{1, "Debt Account Name", []Payment{firstPayment, secondPayment}}
	lastPayment, _ := account.LastPayment()
	if lastPayment != firstPayment {
		t.Errorf("LastPayment with two unordered payments. Expected: %f; Got: %f", firstPayment.Amount, lastPayment.Amount)
	}
}
