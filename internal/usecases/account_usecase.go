package usecases

import (
	"mcandido.com/teste-pismo/external/persistence"
	"mcandido.com/teste-pismo/internal/entities"
)

type AccountUseCase interface {
	CreateAccount(documentNumber string) (*entities.Account, error)
	GetAccount(id int) (*entities.Account, error)
}

type accountUseCase struct {
	accountRepository persistence.AccountRepository
}

func NewAccountUseCase(accountRepository persistence.AccountRepository) *accountUseCase {
	return &accountUseCase{
		accountRepository: accountRepository,
	}
}

func (uc *accountUseCase) CreateAccount(documentNumber string) (*entities.Account, error) {
	// validate input
	acc := &entities.Account{
		DocumentNumber: documentNumber,
	}

	// use repository interfaces to interact with data layer
	acc, err := uc.accountRepository.CreateAccount(acc)
	if err != nil {
		return nil, err
	}

	return acc, nil
}

func (uc *accountUseCase) GetAccount(id int) (*entities.Account, error) {
	// use repository info to get account from layer
	acc, err := uc.accountRepository.FindAccountByID(id)
	if err != nil {
		return nil, err
	}

	return acc, nil
}
