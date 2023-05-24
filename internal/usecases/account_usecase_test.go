package usecases_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"mcandido.com/teste-pismo/internal/entities"
	"mcandido.com/teste-pismo/internal/usecases"
)

// Mock AccountRepository
type MockAccountRepository struct {
	mock.Mock
}

func (m *MockAccountRepository) CreateAccount(account *entities.Account) (*entities.Account, error) {
	args := m.Called(account)
	return args.Get(0).(*entities.Account), args.Error(1)
}

func (m *MockAccountRepository) FindAccountByID(id int) (*entities.Account, error) {
	args := m.Called(id)
	return args.Get(0).(*entities.Account), args.Error(1)
}

// Tests
func TestCreateAccount_Success(t *testing.T) {
	repository := new(MockAccountRepository)
	uc := usecases.NewAccountUseCase(repository)

	givenAccount := &entities.Account{
		DocumentNumber: "123456789",
	}
	wantedAccount := &entities.Account{
		AccountID:      1,
		DocumentNumber: "123456789",
	}
	wanted := *wantedAccount

	repository.
		On("CreateAccount", givenAccount).
		Return(wantedAccount, nil).Once()

	got, err := uc.CreateAccount("123456789")
	assert.NoError(t, err)
	assert.Equal(t, wanted, *got)
}

func TestCreateAccount_RepositoryError(t *testing.T) {
	repository := new(MockAccountRepository)
	uc := usecases.NewAccountUseCase(repository)

	givenAccount := &entities.Account{
		DocumentNumber: "123456789",
	}

	repository.
		On("CreateAccount", givenAccount).
		Return((*entities.Account)(nil), fmt.Errorf("some error")).Once()

	got, err := uc.CreateAccount("123456789")
	assert.Error(t, err)
	assert.Nil(t, got)
}
