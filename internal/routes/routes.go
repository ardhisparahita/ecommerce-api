package routes

import (
	"github.com/ardhisparahita/ecommerce-api/internal/handler"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, authHandler *handler.AuthHandler) {
	api := app.Group("/api/v1")

	auth := api.Group("/auth")

	auth.Post("/register", authHandler.Register)
	auth.Post("/login", authHandler.Login)
}
