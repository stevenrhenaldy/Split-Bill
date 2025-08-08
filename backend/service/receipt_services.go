package service

import (
	"split-bill/backend/dto/request"
	"split-bill/backend/dto/response"

	"github.com/google/uuid"
)

type ReceiptService interface {
	Create(receipt request.CreateReceiptRequest) (response.ReceiptResponse, error)
	FindAll() ([]response.ReceiptResponse, error)
	FindByID(id uuid.UUID) (response.ReceiptResponse, error)
	Update(id uuid.UUID, receipt request.UpdateReceiptRequest) (response.ReceiptResponse, error)
	Delete(id uuid.UUID) error
}
