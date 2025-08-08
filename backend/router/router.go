package router

import (
	"split-bill/backend/controller"

	"github.com/gofiber/fiber/v2"
)

func NewRouter(receiptController *controller.ReceiptController) *fiber.App {
	router := fiber.New()

	router.Route("/receipts", func(router fiber.Router) {
		router.Post("/", receiptController.Create)
		router.Get("/", receiptController.FindAll)
		router.Get("/:id", receiptController.FindByID)
		router.Put("/:id", receiptController.Update)
		router.Delete("/:id", receiptController.Delete)
	})

	return router
}
