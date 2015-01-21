package main

import (
	"os"
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
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	http.ListenAndServe(":" + port, nil)
}
