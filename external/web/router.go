package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"mcandido.com/teste-pismo/internal/interfaces/http/handlers"
)

func SetupRouter(accountsHandler handlers.AccountsHandler, transactionsHandler handlers.TransactionHandler) *gin.Engine {
	r := gin.Default()

	r.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"status": "healthy!",
		})
	})

	accounts := r.Group("/accounts")
	{
		accounts.POST("", accountsHandler.CreateAccount)
		accounts.GET("/:id", accountsHandler.GetAccount)
	}

	transactions := r.Group("/transactions")
	{
		transactions.POST("", transactionsHandler.CreateTransaction)
	}

	return r
}
