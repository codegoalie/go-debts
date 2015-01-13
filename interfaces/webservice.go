package interfaces

import (
	"go-debts/usecases"
	"go-debts/domain"
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
	Debitor(userId int) (domain.Debitor)
}

func (service *WebserviceHandler) ShowAccounts(res http.ResponseWriter, req *http.Request) {
	userId, _ := strconv.Atoi(req.FormValue("userId"))
	accounts, _ := service.UserInteractor.Accounts(userId)

	io.WriteString(res, fmt.Sprintf("%s\n", service.UserInteractor.Debitor(userId).Name))
	for _, account := range accounts {
		io.WriteString(res, fmt.Sprintf("%s\t%.2f\n", account.Name, account.Balance))
	}
}
