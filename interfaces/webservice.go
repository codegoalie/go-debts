package interfaces

import (
	"go-debts/boundaries"
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
	PrepareAccounts(userId int) (boundaries.AccountsOutput, error)
}

func (service *WebserviceHandler) ShowAccounts(res http.ResponseWriter, req *http.Request) {
	userId, _ := strconv.Atoi(req.FormValue("userId"))

	preparedAccounts, _ := service.UserInteractor.PrepareAccounts(userId)


	service.Render.HTML(res, http.StatusOK, "accounts/index", preparedAccounts)
}
