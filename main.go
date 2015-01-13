package main

import (
	"net/http"
	"go-debts/interfaces"
	"go-debts/infrastructure"
	"go-debts/usecases"
)

func main() {
	dbHandler := infrastructure.NewSqliteHandler("/var/tmp/go-debts.sqlite")

	handlers := make(map[string]interfaces.DbHandler)
	handlers["DbUserRepo"] = dbHandler
	handlers["DbDebitorRepo"] = dbHandler
	handlers["DbAccountRepo"] = dbHandler
	handlers["DbPaymentRepo"] = dbHandler

	userInteractor := new(usecases.UserInteractor)
	userInteractor.UserRepository = interfaces.NewDbUserRepo(handlers)
	userInteractor.DebitorRepository = interfaces.NewDbDebitorRepo(handlers)
	userInteractor.AccountRepository = interfaces.NewDbAccountRepo(handlers)
	userInteractor.PaymentRepository = interfaces.NewDbPaymentRepo(handlers)

	webserviceHandler := interfaces.WebserviceHandler{}
	webserviceHandler.UserInteractor = userInteractor

	http.HandleFunc("/accounts", func(res http.ResponseWriter, req *http.Request) {
		webserviceHandler.ShowAccounts(res, req)
	})
	http.ListenAndServe(":8080", nil)
}
