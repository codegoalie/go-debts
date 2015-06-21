package main

type Account struct {
	Id           string  `json:"id"`
	Name         string  `json:"name"`
	InterestRate float64 `json:"interestRate"`
}

type Accounts []Account
