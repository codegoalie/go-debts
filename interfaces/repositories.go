package interfaces

import (
	"go-debts/usecases"
	"go-debts/domain"
	"fmt"
	"time"
)

type DbHandler interface {
	Execute(statement string)
	Query(statement string) Row
}

type Row interface {
	Scan(dest ...interface{})
	Next() bool
}

type DbRepo struct {
	dbHandlers map[string]DbHandler
	dbHandler DbHandler
}

type DbUserRepo DbRepo
type DbDebitorRepo DbRepo
type DbAccountRepo DbRepo
type DbPaymentRepo DbRepo

func NewDbUserRepo(dbHandlers map[string]DbHandler) *DbUserRepo {
	dbUserRepo := new(DbUserRepo)
	dbUserRepo.dbHandlers = dbHandlers
	dbUserRepo.dbHandler = dbHandlers["DbUserRepo"]
	return dbUserRepo
}

func (repo *DbUserRepo) FindById(id int) usecases.User {
	row := repo.dbHandler.Query(fmt.Sprintf("SELECT debitor_id FROM users WHERE id = %d LIMIT 1", id))
	var debitorId int
	row.Next()
	row.Scan(&debitorId)
	debitorRepo := NewDbDebitorRepo(repo.dbHandlers)
	return usecases.User{ID: id, Debitor: debitorRepo.FindById(debitorId)}
}

func NewDbDebitorRepo(dbHandlers map[string]DbHandler) *DbDebitorRepo {
	dbDebitorRepo := new(DbDebitorRepo)
	dbDebitorRepo.dbHandlers = dbHandlers
	dbDebitorRepo.dbHandler = dbHandlers["DbDebitorRepo"]
	return dbDebitorRepo
}

func (repo *DbDebitorRepo) FindById(id int) domain.Debitor {
	row := repo.dbHandler.Query(fmt.Sprintf("SELECT name FROM debitors WHERE id = %d LIMIT 1", id))
	var name string
	row.Next()
	row.Scan(&name)
	debitor := domain.Debitor{ID: id, Name: name}
	var accountId int
	accountRepo := NewDbAccountRepo(repo.dbHandlers)
	row = repo.dbHandler.Query(fmt.Sprintf("SELECT id FROM accounts WHERE debitor_id = %d", id))
	for row.Next() {
		row.Scan(&accountId)
		debitor.Add(accountRepo.FindById(accountId))
	}
	return debitor
}

func NewDbAccountRepo(dbHandlers map[string]DbHandler) *DbAccountRepo {
	dbAccountRepo := new(DbAccountRepo)
	dbAccountRepo.dbHandlers = dbHandlers
	dbAccountRepo.dbHandler = dbHandlers["DbAccountRepo"]
	return dbAccountRepo
}

func (repo *DbAccountRepo) FindById(id int) domain.Account {
	row := repo.dbHandler.Query(fmt.Sprintf("SELECT name FROM accounts WHERE id = %d", id))
	var name string
	row.Next()
	row.Scan(&name)
	account := domain.Account{ID: id, Name: name}
	var paymentId int
	paymentRepo := NewDbPaymentRepo(repo.dbHandlers)
	row = repo.dbHandler.Query(fmt.Sprintf("SELECT id FROM payments WHERE account_id= %d", id))
	for row.Next() {
		row.Scan(&paymentId)
		account.Add(paymentRepo.FindById(paymentId))
	}
	return account
}

func NewDbPaymentRepo(dbHandlers map[string]DbHandler) *DbPaymentRepo {
	dbPaymentRepo := new(DbPaymentRepo)
	dbPaymentRepo.dbHandlers = dbHandlers
	dbPaymentRepo.dbHandler = dbHandlers["DbPaymentRepo"]
	return dbPaymentRepo
}

func (repo *DbPaymentRepo) FindById(id int) domain.Payment {
	row := repo.dbHandler.Query(fmt.Sprintf("SELECT amount, balance, paid_at FROM payments WHERE account_id = %d", id))
	var amount, balance float64
	var paidAt time.Time
	row.Next()
	row.Scan(&amount, &balance, &paidAt)
	return domain.Payment{ID: id, Amount: amount, Balance: balance, Date: paidAt}
}
