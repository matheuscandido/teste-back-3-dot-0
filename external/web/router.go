package web

import (
	"github.com/gin-gonic/gin"
	"mcandido.com/teste-pismo/internal/interfaces/http/handlers"
)

func SetupRouter(accountsHandler handlers.AccountsHandler, transactionsHandler handlers.TransactionHandler) *gin.Engine {
	r := gin.Default()

	accounts := r.Group("/accounts")
	{
		accounts.POST("/", accountsHandler.CreateAccount)
		accounts.GET("/:accountId", accountsHandler.GetAccount)
	}

	transactions := r.Group("/transactions")
	{
		transactions.POST("/", transactionsHandler.CreateTransaction)
	}

	return r
}
