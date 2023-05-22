package main

import (
	"mcandido.com/teste-pismo/external/persistence"
	"mcandido.com/teste-pismo/external/web"
	"mcandido.com/teste-pismo/internal/interfaces/http/handlers"
	"mcandido.com/teste-pismo/internal/usecases"
)

func main() {
	db, _ := persistence.NewDatabase()
	accountsRepository := db
	transactionsRepository := db

	accountsUseCase := usecases.NewAccountUseCase(accountsRepository)
	transactionsUseCase := usecases.NewTransactionUseCase(transactionsRepository)

	accountsHandler := handlers.NewAccountsHandler(accountsUseCase)
	transactionsHandler := handlers.NewTransactionHandler(transactionsUseCase)

	router := web.SetupRouter(accountsHandler, transactionsHandler)
	router.Run(":4000")
}
