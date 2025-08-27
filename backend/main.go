package main

import (
	"fmt"

	"split-bill/backend/config"
	"split-bill/backend/controller"
	"split-bill/backend/model"
	"split-bill/backend/repository"
	"split-bill/backend/router"
	"split-bill/backend/service"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
)

func main() {

	fmt.Println("Starting API server...")

	loadConfig, err := config.LoadConfigFromFile(".")
	if err != nil {
		fmt.Printf("Error loading config: %v\n", err)
		fmt.Println("Reading from Environment Variable")
		loadConfig, err = config.LoadConfigFromEnv()
		if err != nil {
			fmt.Printf("Error loading config from env: %v\n", err)
			return
		}
	}

	// Database
	db := config.ConnectDB(&loadConfig)
	db.AutoMigrate(
		&model.Receipt{},
		&model.ReceiptItem{},
		&model.Currency{},
		&model.User{},
		&model.PaymentInfo{},
		&model.Settlement{},
		&model.Share{},
		&model.ReceiptItemShare{},
	)

	// Init JWT
	jwtConfig := config.NewJwtConfig(loadConfig.JWTSecret, loadConfig.JWTLifetimeHour)

	// Init Repositories
	receiptRepository := repository.NewReceiptRepositoryImpl(db)
	userRepository := repository.NewUserRepositoryImpl(db)

	// Init Validators
	validate := validator.New()

	// Init Services
	receiptService := service.NewReceiptServiceImpl(receiptRepository, validate)
	authService := service.NewAuthServiceImpl(userRepository, validate, jwtConfig)

	// Init Controllers
	receiptController := controller.NewReceiptController(receiptService)
	authController := controller.NewAuthController(authService)

	// Routes
	routes := router.NewRouter(receiptController, authController)

	app := fiber.New()

	app.Mount("/api", routes)

	routes.Get("/healthcheck", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status": "ok",
		})
	})

	app.Listen(":3000")
}
