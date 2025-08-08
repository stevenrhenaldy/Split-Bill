package controller

import (
	"split-bill/backend/dto/request"
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

	receiptResponse, err := controller.ReceiptService.Create(createReceiptRequest)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.Status(fiber.StatusCreated).JSON(receiptResponse)
}

func (controller *ReceiptController) FindAll(ctx *fiber.Ctx) error {
	receipts, err := controller.ReceiptService.FindAll()
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

	receipt, err := controller.ReceiptService.FindByID(parsedUUID)
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

	updatedReceipt, err := controller.ReceiptService.Update(parsedUUID, updateReceiptRequest)
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

	if err := controller.ReceiptService.Delete(parsedUUID); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.Status(fiber.StatusNoContent).JSON(nil)
}
