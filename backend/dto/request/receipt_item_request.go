package request

import "github.com/shopspring/decimal"

type CreateReceiptItemRequest struct {
	DisplayName       string          `json:"name" validate:"required,min=1,max=100"`
	Quantity          int             `json:"quantity" validate:"required,gt=0"`
	TotalItemPrice    decimal.Decimal `json:"total_item_price" validate:"required,decimal"`
	TotalItemDiscount decimal.Decimal `json:"total_item_discount" validate:"omitempty,decimal"`
}

type UpdateReceiptItemRequest struct {
	DisplayName       string          `json:"name" validate:"required,min=1,max=100"`
	Quantity          int             `json:"quantity" validate:"required,gt=0"`
	TotalItemPrice    decimal.Decimal `json:"total_item_price" validate:"required,decimal"`
	TotalItemDiscount decimal.Decimal `json:"total_item_discount" validate:"omitempty,decimal"`
}
