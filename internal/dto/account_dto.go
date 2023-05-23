package dto

type CreateAccountRequest struct {
	DocumentNumber string `json:"document_number" validate:"required"`
}
