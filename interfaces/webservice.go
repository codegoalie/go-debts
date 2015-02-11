package interfaces

import (
	"go-debts/usecases"
	"go-debts/domain"
	"net/http"
	"strconv"
	// "io"
	// "fmt"
	"github.com/unrolled/render"
)


type WebserviceHandler struct {
	UserInteractor UserInteractor
	Render *render.Render
}

type UserInteractor interface {
	Accounts(userId int) ([]usecases.Account, error)
	Debitor(userId int) (domain.Debitor)
}

func (service *WebserviceHandler) ShowAccounts(res http.ResponseWriter, req *http.Request) {
	userId, _ := strconv.Atoi(req.FormValue("userId"))

	accounts, _ := service.UserInteractor.Accounts(userId)
	presenter := AccountsPresenter{ UserName: service.UserInteractor.Debitor(userId).Name,
								    Accounts: accounts }


	service.Render.HTML(res, http.StatusOK, "accounts/index", presenter)
}
