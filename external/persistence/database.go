package persistence

import "mcandido.com/teste-pismo/internal/entities"

type AccountRepository interface {
	CreateAccount(account *entities.Account) (*entities.Account, error)
	FindAccountByID(id int) (*entities.Account, error)
}

type TransactionRepository interface {
	CreateTransaction(transaction *entities.Transaction) (*entities.Transaction, error)
}

type Database struct {
	// db client here
}

func NewDatabase() (*Database, error) {
	// init and configure db connection

	db := &Database{}

	return db, nil
}

func (db *Database) CreateAccount(account *entities.Account) (*entities.Account, error) {
	return nil, nil
}

func (db *Database) FindAccountByID(id int) (*entities.Account, error) {
	return nil, nil
}

func (db *Database) CreateTransaction(transaction *entities.Transaction) (*entities.Transaction, error) {
	return nil, nil
}
