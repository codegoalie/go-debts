package interfaces

import (
	"go-debts/usecases"
	"net/http"
	"strconv"
	"io"
	"fmt"
)

type WebserviceHandler struct {
	UserInteractor UserInteractor
}

type UserInteractor interface {
	Accounts(userId int) ([]usecases.Account, error)
}

func (service *WebserviceHandler) ShowAccounts(res http.ResponseWriter, req *http.Request) {
	userId, _ := strconv.Atoi(req.FormValue("userId"))
	accounts, _ := service.UserInteractor.Accounts(userId)

	for _, account := range accounts {
		io.WriteString(res, fmt.Sprintf("%s\t%f", account.Name, account.Balance))
	}
}
