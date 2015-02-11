package boundaries

import "go-debts/usecases"

type AccountsInput interface {
	Accounts(userId int)
}

type AccountsOutput interface {
	UserName() string
	Accounts() []usecases.Account
}
