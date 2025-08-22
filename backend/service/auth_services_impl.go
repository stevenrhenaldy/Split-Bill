package service

import (
	"split-bill/backend/dto/request"
	"split-bill/backend/dto/response"
	"split-bill/backend/model"
	"split-bill/backend/repository"

	"github.com/go-playground/validator"
)

type AuthServiceImpl struct {
	UserRepository repository.UserRepository
	validate       *validator.Validate
}

func NewAuthServiceImpl(userRepository repository.UserRepository, validate *validator.Validate) AuthService {
	return &AuthServiceImpl{
		UserRepository: userRepository,
		validate:       validate,
	}
}

// Login implements AuthService.
func (s *AuthServiceImpl) Login(request.LoginRequest) (response.AuthResponse, error) {
	panic("unimplemented")
}

// Create implements AuthService.
func (s *AuthServiceImpl) Register(authRequest request.RegisterRequest) error {
	err := s.validate.Struct(authRequest)
	if err != nil {
		return err
	}

	userModel := &model.User{
		Name:              authRequest.Name,
		Username:          authRequest.Username,
		Email:             authRequest.Email,
		Password:          authRequest.Password,
		DefaultCurrencyID: authRequest.DefaultCurrencyID,
		EmailVerifiedAt:   nil,
	}

	_, err = s.UserRepository.Create(userModel)

	if err != nil {
		return err
	}

	return nil
}
