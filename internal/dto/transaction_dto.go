package dto

type CreateTransactionRequest struct {
	AccountID       int     `json:"account_id" validate:"required"`
	OperationTypeID int     `json:"operation_type_id" validate:"required"`
	Amount          float64 `json:"amount" validate:"required"`
}
