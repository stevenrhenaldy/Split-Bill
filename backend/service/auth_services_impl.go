package service

import (
	"split-bill/backend/config"
	"split-bill/backend/dto/request"
	"split-bill/backend/dto/response"
	"split-bill/backend/model"
	"split-bill/backend/repository"

	"github.com/go-playground/validator"
)

type AuthServiceImpl struct {
	UserRepository repository.UserRepository
	validate       *validator.Validate
	jwtConfig      *config.JwtConfig
}

func NewAuthServiceImpl(userRepository repository.UserRepository, validate *validator.Validate, jwtConfig *config.JwtConfig) AuthService {
	return &AuthServiceImpl{
		UserRepository: userRepository,
		validate:       validate,
		jwtConfig:      jwtConfig,
	}
}

// Login implements AuthService.
func (s *AuthServiceImpl) Login(loginRequest request.LoginRequest) (tokenResponse response.TokenResponse, err error) {
	err = s.validate.Struct(loginRequest)
	if err != nil {
		return tokenResponse, err
	}

	user, err := s.UserRepository.FindByUsername(loginRequest.Username)
	if err != nil {
		return tokenResponse, err
	}

	if err := config.CheckPasswordHash(user.Password, loginRequest.Password); err != nil {
		return tokenResponse, err
	}

	// Generate JWT Token
	token, err := s.jwtConfig.GenerateJWT(loginRequest.Username)
	if err != nil {
		return tokenResponse, err
	}

	tokenResponse.Token = token
	return tokenResponse, nil
}

// Create implements AuthService.
func (s *AuthServiceImpl) Register(authRequest request.RegisterRequest) error {
	err := s.validate.Struct(authRequest)
	if err != nil {
		return err
	}

	// Apply Password Hash
	hashed_password, err := config.HashPassword(authRequest.Password)
	if err != nil {
		return err
	}

	userModel := &model.User{
		Name:              authRequest.Name,
		Username:          authRequest.Username,
		Email:             authRequest.Email,
		Password:          hashed_password,
		DefaultCurrencyID: authRequest.DefaultCurrencyID,
		EmailVerifiedAt:   nil,
	}

	_, err = s.UserRepository.Create(userModel)

	if err != nil {
		return err
	}

	return nil
}

// RenewToken implements AuthService.
func (s *AuthServiceImpl) RenewToken() (response.TokenResponse, error) {
	panic("unimplemented")
}

// ForgetPassword implements AuthService.
func (s *AuthServiceImpl) ForgetPassword(request.ForgetPasswordRequest) error {
	panic("unimplemented")
}

// Logout implements AuthService.
func (s *AuthServiceImpl) Logout() error {
	panic("unimplemented")
}

// Me implements AuthService.
func (s *AuthServiceImpl) Me() (response.MeResponse, error) {
	panic("unimplemented")
}

// ResetPassword implements AuthService.
func (s *AuthServiceImpl) ResetPassword(request.ResetPasswordRequest) error {
	panic("unimplemented")
}
