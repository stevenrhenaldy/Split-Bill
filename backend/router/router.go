package router

import (
	"split-bill/backend/controller"

	"github.com/gofiber/fiber/v2"
)

func NewRouter(authMiddleware fiber.Handler, receiptController *controller.ReceiptController, authController *controller.AuthController) *fiber.App {
	router := fiber.New()

	router.Route("/receipts", func(router fiber.Router) {
		router.Post("/", authMiddleware, receiptController.Create)
		router.Get("/", authMiddleware, receiptController.FindAll)
		router.Get("/:id", authMiddleware, receiptController.FindByID)
		router.Put("/:id", authMiddleware, receiptController.Update)
		router.Delete("/:id", authMiddleware, receiptController.Delete)
	})

	router.Route("/auth", func(router fiber.Router) {
		router.Post("/register", authController.Register)
		router.Post("/login", authController.Login)
		router.Post("/renew-token", authController.RenewToken)
		router.Post("/logout", authController.Logout)
		router.Get("/me", authMiddleware, authController.Me)
	})

	return router
}
