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
	FindDebitorByUserId(id int) domain.Debitor
}

type AccountsPresenter struct {
	userName string
	accounts []boundaries.Account
}


func (interactor *UserInteractor) PrepareAccounts(userId int) (boundaries.AccountsOutput, error) {
	debitor := interactor.UserRepository.FindDebitorByUserId(userId)

	accounts := make([]boundaries.Account, len(debitor.Accounts))
	for i, domainAccount := range debitor.Accounts {
		accounts[i] = boundaries.Account{ Name: domainAccount.Name, Balance: domainAccount.CurrentBalance()  }
	}
	return AccountsPresenter{ userName: debitor.Name, accounts: accounts }, nil
}

func (presenter AccountsPresenter) UserName() string {
	return presenter.userName
}

func (presenter AccountsPresenter) Accounts() []boundaries.Account {
	return presenter.accounts
}
