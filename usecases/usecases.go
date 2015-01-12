package usecases

import (
	"go-debts/domain"
)

type UserInteractor struct {
	UserRepository UserRepository
	DebitorRepository domain.DebitorRepository
	AccountRepository domain.AccountRepository
	PaymentRepository domain.PaymentRepository
}

type UserRepository interface {
	FindById(id int) User
}

type User struct {
	ID int
	Debitor domain.Debitor
}

type Account struct {
	ID int
	Name string
	Balance float64
}

func (interactor *UserInteractor) Accounts(userId int) ([]Account, error) {
	user := interactor.UserRepository.FindById(userId)

	accounts := make([]Account, len(user.Debitor.Accounts))
	for i, account := range user.Debitor.Accounts {
		accounts[i] = Account{ID: account.ID, Name: account.Name, Balance: account.CurrentBalance()}
	}
	return accounts, nil
}
