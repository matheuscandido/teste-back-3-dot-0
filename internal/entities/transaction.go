package entities

type Transaction struct {
	TransactionID   int
	AccountID       int
	OperationTypeID int
	Amount          float64
	EventDate       string
}
