package controller

import (
	"split-bill/backend/dto/request"
	"split-bill/backend/service"

	"github.com/gofiber/fiber/v2"
)

type AuthController struct {
	AuthService service.AuthService
}

func NewAuthController(authService service.AuthService) *AuthController {
	return &AuthController{
		AuthService: authService,
	}
}

func (controller *AuthController) Register(ctx *fiber.Ctx) error {
	var registerRequest request.RegisterRequest
	if err := ctx.BodyParser(&registerRequest); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	err := controller.AuthService.Register(registerRequest)
	if err != nil {

		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "User registered successfully"})
}

func (controller *AuthController) Login(ctx *fiber.Ctx) error {
	var loginRequest request.LoginRequest
	if err := ctx.BodyParser(&loginRequest); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	loginResponse, err := controller.AuthService.Login(loginRequest)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.Status(fiber.StatusOK).JSON(loginResponse)
}
