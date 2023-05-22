package usecases

import (
	"mcandido.com/teste-pismo/external/persistence"
	"mcandido.com/teste-pismo/internal/entities"
)

type TransactionUseCase interface {
	CreateTransaction(accountID int, operationTypeID int, amount float64) (*entities.Transaction, error)
}

type transactionUseCase struct {
	transactionRepository persistence.TransactionRepository
}

func NewTransactionUseCase(transactionRepository persistence.TransactionRepository) *transactionUseCase {
	return &transactionUseCase{
		transactionRepository: transactionRepository,
	}
}

func (uc *transactionUseCase) CreateTransaction(accountID int, operationTypeID int, amount float64) (*entities.Transaction, error) {
	// validate input
	transaction := &entities.Transaction{
		AccountID:       accountID,
		OperationTypeID: operationTypeID,
		Amount:          amount,
	}

	// use repository to interface with data layer
	transaction, err := uc.transactionRepository.CreateTransaction(transaction)
	if err != nil {
		return nil, err
	}

	return transaction, nil
}
