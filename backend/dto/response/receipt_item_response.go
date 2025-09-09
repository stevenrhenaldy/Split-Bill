package response

import "github.com/shopspring/decimal"

type ReceiptItemResponse struct {
	ID                string          `json:"id"`
	ReceiptID         string          `json:"receipt_id"`
	DisplayName       string          `json:"display_name"`
	Quantity          int             `json:"quantity"`
	TotalItemPrice    decimal.Decimal `json:"total_item_price"`
	TotalItemDiscount decimal.Decimal `json:"total_item_discount"`
}
