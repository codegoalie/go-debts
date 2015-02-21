package boundaries

type AccountsInput interface {
	Accounts(userId int)
}
