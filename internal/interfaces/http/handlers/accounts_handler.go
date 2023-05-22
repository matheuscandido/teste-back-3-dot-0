package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
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
	// get json data here
	docNumber := "123"

	account, err := h.accountsUseCase.CreateAccount(docNumber)
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
