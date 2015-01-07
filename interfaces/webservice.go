package interfaces

import (
	"go-debts/usecases"
	"net/http"
)

type WebserviceHandler struct {
	UserInteractor UserInteractor
}

type UserInteractor interface {
	Accounts(userId int) ([]usecases.Account, error)
}

func (service *WebserviceHandler) ShowAccounts(res http.ResponseWriter, req *http.Request) {
	// get value from params
	// get accounts from interactor
	// write response
}
