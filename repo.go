package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	// "github.com/nu7hatch/gouuid"
	"log"
	"os"
	"strconv"
)

var db *sql.DB
var logger *log.Logger

func init() {
	logger = log.New(os.Stderr, "Repo: ", 0)

	var err error
	db, err = sql.Open("postgres", "user=debtuser password=password host=localhost dbname=debts sslmode=disable")
	if err != nil {
		log.Fatal("Error opening database: ", err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal("Error pinging database: ", err)
	}
}

func RepoAllAccounts() (loaded Accounts) {
	rows, err := db.Query("SELECT * FROM accounts")
	if err != nil {
		logger.Println("Could not fetch accounts: ", err)
		return Accounts{}
	}
	defer rows.Close()

	accounts := Accounts{}
	var rawInterestRate string
	for rows.Next() {
		account := new(Account)
		err := rows.Scan(&account.Id, &account.Name, &rawInterestRate)
		if err != nil {
			logger.Println("Error scanning accounts: ", err)
			continue
		}
		account.InterestRate, err = strconv.ParseFloat(rawInterestRate, 64)
		if err != nil {
			logger.Println("Error parsing interest rate: ", err)
			continue
		}
		accounts = append(accounts, *account)
	}

	return accounts
}
