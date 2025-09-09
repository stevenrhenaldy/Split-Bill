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
func (r *ReceiptServiceImpl) Create(user *model.User, receipt request.CreateReceiptRequest) (response.ReceiptResponse, error) {
	err := r.validate.Struct(receipt)
	if err != nil {
		return response.ReceiptResponse{}, err
	}

	receiptModel := &model.Receipt{
		Name:        receipt.Name,
		UserID:      user.ID,
		Description: receipt.Description,
	}
	receiptModel, err = r.ReceiptRepository.Create(user, receiptModel)
	if err != nil {
		return response.ReceiptResponse{}, err
	}
	return response.ReceiptResponse{
		ID:          receiptModel.ID.String(),
		Name:        receiptModel.Name,
		Description: receiptModel.Description,
		UserID:      receiptModel.UserID.String(),
		CreatedAt:   receiptModel.CreatedAt.String(),
		UpdatedAt:   receiptModel.UpdatedAt.String(),
	}, nil
}

// Delete implements ReceiptService.
func (r *ReceiptServiceImpl) Delete(user *model.User, id uuid.UUID) error {
	return r.ReceiptRepository.Delete(user, id.String())
}

// FindAll implements ReceiptService.
func (r *ReceiptServiceImpl) FindAll(user *model.User) ([]response.ReceiptResponse, error) {
	receipts, err := r.ReceiptRepository.FindAll(user)
	if err != nil {
		return []response.ReceiptResponse{}, err
	}

	var responses []response.ReceiptResponse
	for _, receipt := range receipts {
		responses = append(responses, response.ReceiptResponse{
			ID:          receipt.ID.String(),
			Name:        receipt.Name,
			Description: receipt.Description,
			UserID:      receipt.UserID.String(),
			CreatedAt:   receipt.CreatedAt.String(),
			UpdatedAt:   receipt.UpdatedAt.String(),
		})
	}
	return responses, nil
}

// FindByID implements ReceiptService.
func (r *ReceiptServiceImpl) FindByID(user *model.User, id uuid.UUID) (response.ReceiptResponse, error) {
	receipt, err := r.ReceiptRepository.FindByID(user, id.String())
	if err != nil {
		return response.ReceiptResponse{}, err
	}
	return response.ReceiptResponse{
		ID:          receipt.ID.String(),
		Name:        receipt.Name,
		Description: receipt.Description,
		UserID:      receipt.UserID.String(),
		CreatedAt:   receipt.CreatedAt.String(),
		UpdatedAt:   receipt.UpdatedAt.String(),
	}, nil
}

// Update implements ReceiptService.
func (r *ReceiptServiceImpl) Update(user *model.User, id uuid.UUID, receipt request.UpdateReceiptRequest) (response.ReceiptResponse, error) {
	err := r.validate.Struct(receipt)
	if err != nil {
		return response.ReceiptResponse{}, err
	}

	receiptModel := &model.Receipt{
		ID:          id,
		Name:        receipt.Name,
		Description: receipt.Description,
	}

	receiptModel, err = r.ReceiptRepository.Update(user, receiptModel)
	if err != nil {
		return response.ReceiptResponse{}, err
	}

	return response.ReceiptResponse{
		ID:          receiptModel.ID.String(),
		Name:        receiptModel.Name,
		Description: receiptModel.Description,
		UserID:      receiptModel.UserID.String(),
		CreatedAt:   receiptModel.CreatedAt.String(),
		UpdatedAt:   receiptModel.UpdatedAt.String(),
	}, nil
}
