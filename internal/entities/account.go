package entities

type Account struct {
	AccountID      int    `json:"account_id,omitempty"`
	DocumentNumber string `json:"document_number,omitempty"`
}

type OperationType struct {
	OperationTypeID int
	Description     string
}
