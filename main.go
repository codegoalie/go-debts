package main

import (
	_ "github.com/mattn/go-sqlite3"
	"github.com/unrolled/render"
	"go-debts/infrastructure"
	"go-debts/interfaces"
	"net/http"
	"os"
)

func main() {
	r := render.New(render.Options{Layout: "layout"})

	dbHandler := infrastructure.NewSqliteHandler("/var/tmp/go-debts.sqlite")

	accountController := accountController{r: r, handler: dbHandler}

	http.HandleFunc("/accounts", accountController.index)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	http.ListenAndServe(":"+port, nil)
}

type AccountsListingInput interface {
	fetchAccountsForUser(userId int) accountsViewModel
}

type AccountsListingUseCase struct {
	handler        interfaces.DbHandler
	userGateway    userGateway
	accountGateway accountGateway
}

func (usecase AccountsListingUseCase) fetchAccountsForUser(userId int) accountsViewModel {
	debitor := usecase.userGateway.fetchDebitorByUserId(userId)
	return accountsViewModel{UserName: debitor.name,
		Accounts: usecase.accountGateway.fetchAccountsByDebitorId(debitor.id)}
}

type accountGateway interface {
	fetchAccountsByDebitorId(debitorId int) []account
	fetchAccountByID(accountID int) account
}

type userGateway interface {
	fetchDebitorByUserId(userId int) debitor
}
