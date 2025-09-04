package service

import (
	"split-bill/backend/dto/request"
	"split-bill/backend/dto/response"
	"split-bill/backend/model"

	"github.com/gofiber/fiber/v2"
)

type AuthService interface {
	Register(request.RegisterRequest) error
	Login(*fiber.Ctx, request.LoginRequest) error
	Logout(*fiber.Ctx) error
	RenewToken(*fiber.Ctx) error
	Me(model.User) (response.MeResponse, error)
	ForgetPassword(model.User, request.ForgetPasswordRequest) error
	ResetPassword(model.User, request.ResetPasswordRequest) error
}
