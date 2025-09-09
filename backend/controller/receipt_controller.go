package controller

import (
	"split-bill/backend/dto/request"
	"split-bill/backend/model"
	"split-bill/backend/service"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type ReceiptController struct {
	ReceiptService service.ReceiptService
}

func NewReceiptController(receiptService service.ReceiptService) *ReceiptController {
	return &ReceiptController{
		ReceiptService: receiptService,
	}
}

func (controller *ReceiptController) Create(ctx *fiber.Ctx) error {
	var createReceiptRequest request.CreateReceiptRequest
	if err := ctx.BodyParser(&createReceiptRequest); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	user := ctx.Locals("user").(*model.User)
	receiptResponse, err := controller.ReceiptService.Create(user, createReceiptRequest)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.Status(fiber.StatusCreated).JSON(receiptResponse)
}

func (controller *ReceiptController) FindAll(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(*model.User)
	receipts, err := controller.ReceiptService.FindAll(user)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.Status(fiber.StatusOK).JSON(receipts)
}

func (controller *ReceiptController) FindByID(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	parsedUUID, err := uuid.Parse(id)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid UUID"})
	}

	user := ctx.Locals("user").(*model.User)
	receipt, err := controller.ReceiptService.FindByID(user, parsedUUID)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.Status(fiber.StatusOK).JSON(receipt)
}

func (controller *ReceiptController) Update(ctx *fiber.Ctx) error {
	var updateReceiptRequest request.UpdateReceiptRequest
	if err := ctx.BodyParser(&updateReceiptRequest); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	id := ctx.Params("id")

	parsedUUID, err := uuid.Parse(id)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid UUID"})
	}

	user := ctx.Locals("user").(*model.User)
	updatedReceipt, err := controller.ReceiptService.Update(user, parsedUUID, updateReceiptRequest)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.Status(fiber.StatusOK).JSON(updatedReceipt)
}

func (controller *ReceiptController) Delete(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	parsedUUID, err := uuid.Parse(id)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid UUID"})
	}

	user := ctx.Locals("user").(*model.User)
	if err := controller.ReceiptService.Delete(user, parsedUUID); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.Status(fiber.StatusNoContent).JSON(nil)
}
