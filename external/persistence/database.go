package persistence

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	"mcandido.com/teste-pismo/internal/entities"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

type AccountRepository interface {
	CreateAccount(account *entities.Account) (*entities.Account, error)
	FindAccountByID(id int) (*entities.Account, error)
}

type TransactionRepository interface {
	CreateTransaction(transaction *entities.Transaction) error
}

var (
	host     = os.Getenv("DB_HOST")
	port     = os.Getenv("DB_PORT")
	user     = os.Getenv("DB_USER")
	password = os.Getenv("DB_PASS")
	dbname   = os.Getenv("DB_NAME")
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

	// ping db
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	// run migrations
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return nil, err
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://../migrations",
		"postgres", driver)
	if err != nil {
		return nil, err
	}

	m.Up()

	// wrap up
	database := &Database{
		DB: db,
	}

	return database, nil
}

func (db *Database) CreateAccount(account *entities.Account) (*entities.Account, error) {
	_, err := db.DB.Exec(fmt.Sprintf("INSERT INTO accounts (document_number) values ('%s')", account.DocumentNumber))
	if err != nil {
		return nil, err
	}

	newAcc := &entities.Account{}

	err = db.DB.
		QueryRow("SELECT id, document_number FROM accounts where document_number = $1", account.DocumentNumber).
		Scan(&newAcc.AccountID, &newAcc.DocumentNumber)
	if err != nil {
		return nil, err
	}

	return newAcc, nil
}

func (db *Database) FindAccountByID(id int) (*entities.Account, error) {
	account := &entities.Account{}

	err := db.DB.
		QueryRow("SELECT id, document_number FROM accounts where id = $1", id).
		Scan(&account.AccountID, &account.DocumentNumber)
	if err != nil {
		return nil, err
	}

	return account, nil
}

func (db *Database) CreateTransaction(transaction *entities.Transaction) error {
	stm := `INSERT INTO transactions (account_id, operation_id, amount, event_date) values ($1, $2, $3, $4)`
	_, err := db.DB.Exec(stm, transaction.AccountID, transaction.OperationTypeID, transaction.Amount, time.Now().UTC().Format(time.RFC3339))
	if err != nil {
		return err
	}

	return nil
}
