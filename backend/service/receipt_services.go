package service

import (
	"split-bill/backend/dto/request"
	"split-bill/backend/dto/response"
	"split-bill/backend/model"

	"github.com/google/uuid"
)

type ReceiptService interface {
	Create(*model.User, request.CreateReceiptRequest) (response.ReceiptResponse, error)
	FindAll(*model.User) ([]response.ReceiptResponse, error)
	FindByID(*model.User, uuid.UUID) (response.ReceiptResponse, error)
	Update(*model.User, uuid.UUID, request.UpdateReceiptRequest) (response.ReceiptResponse, error)
	Delete(*model.User, uuid.UUID) error
}
