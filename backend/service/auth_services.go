package service

import (
	"split-bill/backend/dto/request"
	"split-bill/backend/dto/response"

	"github.com/gofiber/fiber/v2"
)

type AuthService interface {
	Register(request.RegisterRequest) error
	Login(*fiber.Ctx, request.LoginRequest) error
	Logout(*fiber.Ctx) error
	RenewToken(*fiber.Ctx) error
	Me() (response.MeResponse, error)
	ForgetPassword(request.ForgetPasswordRequest) error
	ResetPassword(request.ResetPasswordRequest) error
}
