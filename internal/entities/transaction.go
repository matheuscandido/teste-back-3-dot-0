package entities

type Transaction struct {
	TransactionID   int     `json:"transaction_id,omitempty"`
	AccountID       int     `json:"account_id,omitempty"`
	OperationTypeID int     `json:"operation_type_id,omitempty"`
	Amount          float64 `json:"amount,omitempty"`
	EventDate       string  `json:"event_date,omitempty"`
}
