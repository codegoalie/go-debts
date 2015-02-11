package usecases

import (
	"go-debts/domain"
	"go-debts/boundaries"
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

type AccountsPresenter struct {
	UserName string
	Accounts []Account
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

func (interactor *UserInteractor) Accounts(userId int) (boundaries.AccountsOutput, error) {
	user := interactor.UserRepository.FindById(userId)

	accounts := make([]Account, len(user.Debitor.Accounts))
	for i, account := range user.Debitor.Accounts {
		accounts[i] = Account{ID: account.ID, Name: account.Name, Balance: account.CurrentBalance()}
	}
	return accounts, nil
}

func (interactor *UserInteractor) Debitor(userId int) (domain.Debitor) {
	user := interactor.UserRepository.FindById(userId)
	return user.Debitor
}

func (presenter *AccountsPresenter) UserName() {
	return presenter.UserName
}
