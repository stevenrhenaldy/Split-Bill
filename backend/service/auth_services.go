package service

import (
	"split-bill/backend/dto/request"
	"split-bill/backend/dto/response"
)

type AuthService interface {
	Register(request.RegisterRequest) error
	Login(request.LoginRequest) (response.AuthResponse, error)
}
