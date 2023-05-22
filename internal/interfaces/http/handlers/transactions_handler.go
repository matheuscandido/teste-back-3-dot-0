package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"mcandido.com/teste-pismo/internal/usecases"
)

type TransactionHandler interface {
	CreateTransaction(c *gin.Context)
}

type transactionHandler struct {
	transactionsUseCase usecases.TransactionUseCase
}

func NewTransactionHandler(transactionUseCase usecases.TransactionUseCase) TransactionHandler {
	return &transactionHandler{
		transactionsUseCase: transactionUseCase,
	}
}

func (h *transactionHandler) CreateTransaction(c *gin.Context) {
	// get data here
	accountID := 1
	operationTypeID := 1
	amount := 1.12

	// create transaction
	transaction, err := h.transactionsUseCase.CreateTransaction(accountID, operationTypeID, amount)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create transaction",
		})
		return
	}

	c.JSON(http.StatusOK, transaction)
}
