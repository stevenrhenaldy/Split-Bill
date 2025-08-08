package request

type CreateReceiptRequest struct {
	Name        string `json:"name" validate:"required,min=2,max=100"`
	Description string `json:"description" validate:"omitempty,max=500"`
}

type UpdateReceiptRequest struct {
	Name        string `json:"name" validate:"min=2,max=100"`
	Description string `json:"description" validate:"omitempty,max=500"`
}
