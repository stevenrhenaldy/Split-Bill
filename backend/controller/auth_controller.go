package controller

import (
	"split-bill/backend/dto/request"
	"split-bill/backend/model"
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

	err := controller.AuthService.Login(ctx, loginRequest)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Login successful"})
}

func (controller *AuthController) RenewToken(ctx *fiber.Ctx) error {
	err := controller.AuthService.RenewToken(ctx)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Token renewed successfully"})
}

func (controller *AuthController) Logout(ctx *fiber.Ctx) error {
	err := controller.AuthService.Logout(ctx)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Logout successful"})
}

func (controller *AuthController) Me(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(*model.User)
	response, err := controller.AuthService.Me(*user)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.Status(fiber.StatusOK).JSON(response)
}
