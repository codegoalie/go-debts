package main

import (
	"github.com/unrolled/render"
	"go-debts/infrastructure"
	"go-debts/interfaces"
	"go-debts/usecases"
	"net/http"
	"os"
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

	r := render.New()

	webserviceHandler := interfaces.WebserviceHandler{}
	webserviceHandler.UserInteractor = userInteractor
	webserviceHandler.Render = r

	http.HandleFunc("/accounts", func(res http.ResponseWriter, req *http.Request) {
		webserviceHandler.ShowAccounts(res, req)
	})
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	http.ListenAndServe(":"+port, nil)
}
