package service

import (
	"split-bill/backend/dto/request"
	"split-bill/backend/dto/response"
)

type AuthService interface {
	Register(request.RegisterRequest) error
	Login(request.LoginRequest) (response.TokenResponse, error)
	Logout() error
	RenewToken() (response.TokenResponse, error)
	Me() (response.MeResponse, error)
	ForgetPassword(request.ForgetPasswordRequest) error
	ResetPassword(request.ResetPasswordRequest) error
}
