package main

import (
	"github.com/unrolled/render"
	"net/http"
	"os"
	"strconv"
)

func main() {
	r := render.New(render.Options{Layout: "layout"})

	accountController := accountController{r: r}

	http.HandleFunc("/accounts", accountController.index)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	http.ListenAndServe(":"+port, nil)
}

type accountController struct {
	r *render.Render
}

func (controller accountController) index(res http.ResponseWriter, req *http.Request) {
	userId, _ := strconv.Atoi(req.FormValue("userId"))

	accountsListing := AccountsListingUseCase{gateway: staticUserGateway{}}
	viewModel := accountsListing.fetchAccountsForUser(userId)

	controller.r.HTML(res, http.StatusOK, "accounts/index", viewModel)
}

type AccountsListingInput interface {
	fetchAccountsForUser(userId int) accountViewModel
}

type AccountsListingUseCase struct {
	gateway userGateway
}

func (usecase AccountsListingUseCase) fetchAccountsForUser(userId int) accountViewModel {
	return accountViewModel{UserName: usecase.gateway.fetchUserById(userId).userName,
		Accounts: fetchAccountByUserId(userId)}
}

type accountViewModel struct {
	UserName string
	Accounts []account
}

type account struct {
	ID      int
	Name    string
	Balance float64
}

type user struct {
	ID       int
	email    string
	userName string
}

func fetchAccountByUserId(userId int) []account {
	return []account{
		account{ID: 3, Name: "Bank of America", Balance: 54.55},
		account{ID: 4, Name: "Citi Card", Balance: 652.74}}
}

type userGateway interface {
	fetchUserById(ID int) user
}

type staticUserGateway struct {
}

func (gateway staticUserGateway) fetchUserById(ID int) user {
	return user{ID: 1, email: "chrismar035@gmail.com", userName: "Chris Marshall"}
}
