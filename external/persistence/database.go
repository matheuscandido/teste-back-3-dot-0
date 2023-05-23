package persistence

import (
	"database/sql"
	"fmt"
	"time"

	"mcandido.com/teste-pismo/internal/entities"

	_ "github.com/lib/pq"
)

type AccountRepository interface {
	CreateAccount(account *entities.Account) (*entities.Account, error)
	FindAccountByID(id int) (*entities.Account, error)
}

type TransactionRepository interface {
	CreateTransaction(transaction *entities.Transaction) error
}

const (
	host     = "localhost"
	port     = 5432
	user     = "yourusername"
	password = "yourpassword"
	dbname   = "yourdbname"
)

type Database struct {
	DB *sql.DB
}

func NewDatabase() (*Database, error) {
	connectionString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	database := &Database{
		DB: db,
	}

	return database, nil
}

func (db *Database) CreateAccount(account *entities.Account) (*entities.Account, error) {
	stm := `INSERT INTO accounts (document_number) values ($1)`
	_, err := db.DB.Exec(stm, account.DocumentNumber)
	if err != nil {
		return nil, err
	}

	newAcc := entities.Account{}

	err = db.DB.QueryRow("SELECT id, document_number FROM accounts where document_number = ?", account.DocumentNumber).
		Scan(newAcc.AccountID, newAcc.DocumentNumber)
	if err != nil {
		return nil, err
	}

	return &newAcc, nil
}

func (db *Database) FindAccountByID(id int) (*entities.Account, error) {
	account := &entities.Account{}

	err := db.DB.
		QueryRow("SELECT id, document_number FROM accounts where id = ?", id).
		Scan(account.AccountID, account.DocumentNumber)
	if err != nil {
		return nil, err
	}

	return account, nil
}

func (db *Database) CreateTransaction(transaction *entities.Transaction) error {
	stm := `INSERT INTO transactions (account_id, operation_type_id, amount, event_date) values ($1, $2, $3, $4)`
	_, err := db.DB.Exec(stm, transaction.AccountID, transaction.OperationTypeID, transaction.Amount, time.Now().UTC().Format(time.RFC3339))
	if err != nil {
		return err
	}

	return nil
}
