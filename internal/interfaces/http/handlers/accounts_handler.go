package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"mcandido.com/teste-pismo/internal/dto"
	"mcandido.com/teste-pismo/internal/usecases"
)

type AccountsHandler interface {
	CreateAccount(c *gin.Context)
	GetAccount(c *gin.Context)
}

type accountsHandler struct {
	accountsUseCase usecases.AccountUseCase
}

func NewAccountsHandler(accountsUseCase usecases.AccountUseCase) AccountsHandler {
	return &accountsHandler{
		accountsUseCase: accountsUseCase,
	}
}

func (h *accountsHandler) CreateAccount(c *gin.Context) {
	var accountBody dto.CreateAccountRequest
	if err := c.ShouldBindJSON(&accountBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request payload",
		})
		return
	}

	if err := validator.New().Struct(accountBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":  "Invalid request body",
			"fields": err.(validator.ValidationErrors),
		})
		return
	}

	account, err := h.accountsUseCase.CreateAccount(accountBody.DocumentNumber)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create account",
		})
		return
	}

	c.JSON(http.StatusOK, account)
}

func (h *accountsHandler) GetAccount(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid account ID",
		})
		return
	}

	acc, err := h.accountsUseCase.GetAccount(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Account not found",
		})
		return
	}

	c.JSON(http.StatusOK, acc)
}
