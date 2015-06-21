package main

import (
	"encoding/json"
	"net/http"
)

func AccountIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json;cahrset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(RepoAllAccounts()); err != nil {
		panic(err)
	}
}
