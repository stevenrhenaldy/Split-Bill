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

// func (controller *UserController) FindAll(ctx *fiber.Ctx) error {
// 	users, err := controller.UserService.FindAll()
// 	if err != nil {
// 		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
// 	}
// 	return ctx.Status(fiber.StatusOK).JSON(users)
// }

// func (controller *UserController) FindByID(ctx *fiber.Ctx) error {
// 	id := ctx.Params("id")
// 	parsedUUID, err := uuid.Parse(id)
// 	if err != nil {
// 		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid UUID"})
// 	}

// 	user, err := controller.UserService.FindByID(parsedUUID)
// 	if err != nil {
// 		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
// 	}
// 	return ctx.Status(fiber.StatusOK).JSON(user)
// }

// func (controller *UserController) Update(ctx *fiber.Ctx) error {
// 	var updateUserRequest request.UpdateUserRequest
// 	if err := ctx.BodyParser(&updateUserRequest); err != nil {
// 		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
// 	}

// 	id := ctx.Params("id")

// 	parsedUUID, err := uuid.Parse(id)
// 	if err != nil {
// 		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid UUID"})
// 	}

// 	updatedUser, err := controller.UserService.Update(parsedUUID, updateUserRequest)
// 	if err != nil {
// 		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
// 	}

// 	return ctx.Status(fiber.StatusOK).JSON(updatedUser)
// }

// func (controller *UserController) Delete(ctx *fiber.Ctx) error {
// 	id := ctx.Params("id")
// 	parsedUUID, err := uuid.Parse(id)
// 	if err != nil {
// 		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid UUID"})
// 	}

// 	if err := controller.UserService.Delete(parsedUUID); err != nil {
// 		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
// 	}

// 	return ctx.Status(fiber.StatusNoContent).JSON(nil)
// }
