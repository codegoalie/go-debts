package main

import (
	_ "github.com/mattn/go-sqlite3"
	"github.com/unrolled/render"
	"go-debts/interfaces"
	"net/http"
	"strconv"
)

type accountController struct {
	r       *render.Render
	handler interfaces.DbHandler
}

type accountsViewModel struct {
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

type payment struct {
	Balance float64
	Amount  float64
	PaidAt  string
}

func (controller accountController) index(res http.ResponseWriter, req *http.Request) {
	userId, _ := strconv.Atoi(req.FormValue("userId"))

	accountsListing := AccountsListingUseCase{userGateway: dbUserGateway{handler: controller.handler}, accountGateway: dbAccountGateway{handler: controller.handler}}
	viewModel := accountsListing.fetchAccountsForUser(userId)

	controller.r.HTML(res, http.StatusOK, "accounts/index", viewModel)
}
