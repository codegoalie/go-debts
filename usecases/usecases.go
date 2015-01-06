package usecases

import (
	"go-debts/domain"
)

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

type AccountInteractor struct {
	UserRepository UserRepository
}

func (interactor *AccountInteractor) Accounts(userId int) ([]Account, error) {
	user := interactor.UserRepository.FindById(userId)

	accounts := make([]Account, len(user.Debitor.Accounts))
	for i, account := range user.Debitor.Accounts {
		accounts[i] = Account{ID: account.ID, Name: account.Name, Balance: account.CurrentBalance()}
	}
	return accounts, nil
}
