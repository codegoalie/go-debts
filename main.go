package main

import (
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"github.com/unrolled/render"
	"go-debts/infrastructure"
	"go-debts/interfaces"
	"net/http"
	"os"
	"strconv"
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

type accountController struct {
	r       *render.Render
	handler interfaces.DbHandler
}

func (controller accountController) index(res http.ResponseWriter, req *http.Request) {
	userId, _ := strconv.Atoi(req.FormValue("userId"))

	accountsListing := AccountsListingUseCase{userGateway: dbUserGateway{handler: controller.handler}, accountGateway: dbAccountGateway{handler: controller.handler}}
	viewModel := accountsListing.fetchAccountsForUser(userId)

	controller.r.HTML(res, http.StatusOK, "accounts/index", viewModel)
}

type AccountsListingInput interface {
	fetchAccountsForUser(userId int) accountViewModel
}

type AccountsListingUseCase struct {
	handler        interfaces.DbHandler
	userGateway    userGateway
	accountGateway accountGateway
}

func (usecase AccountsListingUseCase) fetchAccountsForUser(userId int) accountViewModel {
	debitor := usecase.userGateway.fetchDebitorByUserId(userId)
	return accountViewModel{UserName: debitor.name,
		Accounts: usecase.accountGateway.fetchAccountsByDebitorId(debitor.id)}
}

type accountViewModel struct {
	UserName string
	Accounts []account
}

type account struct {
	ID      int
	Name    string
	Balance float64
}

type debitor struct {
	id   int
	name string
}

type accountGateway interface {
	fetchAccountsByDebitorId(debitorId int) []account
}

type staticAccountGateway struct {
}

func (gateway staticAccountGateway) fetchAccountsByDebitorId(debitorId int) []account {
	return []account{
		account{ID: 3, Name: "Bank of America", Balance: 54.25},
		account{ID: 4, Name: "Citi Card", Balance: 652.74}}
}

type dbAccountGateway struct {
	handler interfaces.DbHandler
}

func (gateway dbAccountGateway) fetchAccountsByDebitorId(debitorId int) []account {
	row := gateway.handler.Query(fmt.Sprintf("SELECT id, name FROM accounts WHERE debitor_id = %d", debitorId))
	var accountId int
	var name string
	var balance float64
	var accounts []account
	for row.Next() {
		row.Scan(&accountId, &name)
		paymentRow := gateway.handler.Query(fmt.Sprintf("SELECT balance FROM payments WHERE account_id = %d ORDER BY paid_at DESC LIMIT 1", accountId))
		paymentRow.Next()
		paymentRow.Scan(&balance)
		fmt.Println("Balance: ", balance)
		accounts = append(accounts, account{ID: accountId, Name: name, Balance: balance})
	}
	return accounts
}

type userGateway interface {
	fetchDebitorByUserId(userId int) debitor
}

type staticUserGateway struct {
}

func (gateway staticUserGateway) fetchDebitorByUserId(userId int) debitor {
	return debitor{id: userId, name: "Chris Marshall"}
}

type dbUserGateway struct {
	handler interfaces.DbHandler
}

func (gateway dbUserGateway) fetchDebitorByUserId(userId int) debitor {
	row := gateway.handler.Query(fmt.Sprintf("SELECT debitor_id FROM users WHERE id = %d LIMIT 1", userId))
	var debitorId int
	row.Next()
	row.Scan(&debitorId)
	row = gateway.handler.Query(fmt.Sprintf("SELECT name FROM debitors WHERE id = %d LIMIT 1", debitorId))
	var name string
	row.Next()
	row.Scan(&name)
	return debitor{id: debitorId, name: name}
}
