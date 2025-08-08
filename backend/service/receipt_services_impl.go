package service

import (
	"split-bill/backend/dto/request"
	"split-bill/backend/dto/response"
	"split-bill/backend/model"
	"split-bill/backend/repository"

	"github.com/go-playground/validator"
	"github.com/google/uuid"
)

type ReceiptServiceImpl struct {
	ReceiptRepository repository.ReceiptRepository
	validate          *validator.Validate
}

func NewReceiptServiceImpl(receiptRepository repository.ReceiptRepository, validate *validator.Validate) ReceiptService {
	return &ReceiptServiceImpl{
		ReceiptRepository: receiptRepository,
		validate:          validate,
	}
}

// Create implements ReceiptService.
func (r *ReceiptServiceImpl) Create(receipt request.CreateReceiptRequest) (response.ReceiptResponse, error) {
	err := r.validate.Struct(receipt)
	if err != nil {
		return response.ReceiptResponse{}, err
	}

	receiptModel := &model.Receipt{
		Name:        receipt.Name,
		Description: receipt.Description,
	}
	receiptModel, err = r.ReceiptRepository.Create(receiptModel)
	if err != nil {
		return response.ReceiptResponse{}, err
	}
	return response.ReceiptResponse{
		ID:          receiptModel.ID.String(),
		Name:        receiptModel.Name,
		Description: receiptModel.Description,
		CreatedAt:   receiptModel.CreatedAt.String(),
		UpdatedAt:   receiptModel.UpdatedAt.String(),
	}, nil
}

// Delete implements ReceiptService.
func (r *ReceiptServiceImpl) Delete(id uuid.UUID) error {
	return r.ReceiptRepository.Delete(id.String())
}

// FindAll implements ReceiptService.
func (r *ReceiptServiceImpl) FindAll() ([]response.ReceiptResponse, error) {
	receipts, err := r.ReceiptRepository.FindAll()
	if err != nil {
		return []response.ReceiptResponse{}, err
	}

	var responses []response.ReceiptResponse
	for _, receipt := range receipts {
		responses = append(responses, response.ReceiptResponse{
			ID:          receipt.ID.String(),
			Name:        receipt.Name,
			Description: receipt.Description,
			CreatedAt:   receipt.CreatedAt.String(),
			UpdatedAt:   receipt.UpdatedAt.String(),
		})
	}
	return responses, nil
}

// FindByID implements ReceiptService.
func (r *ReceiptServiceImpl) FindByID(id uuid.UUID) (response.ReceiptResponse, error) {
	receipt, err := r.ReceiptRepository.FindByID(id.String())
	if err != nil {
		return response.ReceiptResponse{}, err
	}
	return response.ReceiptResponse{
		ID:          receipt.ID.String(),
		Name:        receipt.Name,
		Description: receipt.Description,
		CreatedAt:   receipt.CreatedAt.String(),
		UpdatedAt:   receipt.UpdatedAt.String(),
	}, nil
}

// Update implements ReceiptService.
func (r *ReceiptServiceImpl) Update(id uuid.UUID, receipt request.UpdateReceiptRequest) (response.ReceiptResponse, error) {
	err := r.validate.Struct(receipt)
	if err != nil {
		return response.ReceiptResponse{}, err
	}

	receiptModel := &model.Receipt{
		ID:          id,
		Name:        receipt.Name,
		Description: receipt.Description,
	}

	receiptModel, err = r.ReceiptRepository.Update(receiptModel)
	if err != nil {
		return response.ReceiptResponse{}, err
	}

	return response.ReceiptResponse{
		ID:          receiptModel.ID.String(),
		Name:        receiptModel.Name,
		Description: receiptModel.Description,
		CreatedAt:   receiptModel.CreatedAt.String(),
		UpdatedAt:   receiptModel.UpdatedAt.String(),
	}, nil
}
