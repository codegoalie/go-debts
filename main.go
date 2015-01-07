package main

import (
	"net/http"
	"go-debts/interfaces"
)

func main() {
	// infrastructure.NewSqliteHandler("/var/tmp/production.sqlite")


	webserviceHandler := interfaces.WebserviceHandler{}

	http.HandleFunc("/orders", func(res http.ResponseWriter, req *http.Request) {
		webserviceHandler.ShowAccounts(res, req)
	})
	http.ListenAndServe(":8080", nil)
}
