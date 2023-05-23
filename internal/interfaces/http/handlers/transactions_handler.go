package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"mcandido.com/teste-pismo/internal/dto"
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
	var transactionBody dto.CreateTransactionRequest
	if err := c.ShouldBindJSON(&transactionBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request payload",
		})
		return
	}

	if err := validator.New().Struct(transactionBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":  "Invalid request body",
			"fields": err.(validator.ValidationErrors),
		})
		return
	}

	err := h.transactionsUseCase.CreateTransaction(transactionBody.AccountID, transactionBody.OperationTypeID, transactionBody.Amount)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create transaction",
		})
		return
	}

	c.Status(http.StatusOK)
}
