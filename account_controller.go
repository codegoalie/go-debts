package main

import (
	"github.com/gorilla/mux"
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

type accountViewModel struct {
	Account  account
	Payments []payment
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

func (controller accountController) show(res http.ResponseWriter, req *http.Request) {
	accountID, _ := strconv.Atoi(mux.Vars(req)["id"])

	accountDetail := AccountDetailUseCase{accountGateway: dbAccountGateway{handler: controller.handler}, paymentGateway: dbPaymentGateway{handler: controller.handler}}
	viewModel := accountDetail.fetchAccountDetails(accountID)

	controller.r.HTML(res, http.StatusOK, "accounts/show", viewModel)
}
