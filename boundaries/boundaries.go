package boundaries

type AccountsInput interface {
	Accounts(userId int)
}

type AccountsOutput interface {
	UserName() string
	Accounts() []Account
}

type Account struct {
	ID int
	Name string
	Balance float64
}
