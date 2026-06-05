package routes

import (
	"github.com/ardhisparahita/ecommerce-api/internal/handler"
	"github.com/ardhisparahita/ecommerce-api/internal/middleware"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, authHandler *handler.AuthHandler, categoryHandler *handler.CategoryHandler, productHandler *handler.ProductHandler, addressHandler *handler.AddressHandler, cartHandler *handler.CartHandler) {
	api := app.Group("/api/v1")

	auth := api.Group("/auth")

	auth.Post("/register", authHandler.Register)
	auth.Post("/login", authHandler.Login)

	category := api.Group("/categories", middleware.JWT())
	category.Post("/", categoryHandler.Create)
	category.Get("/", categoryHandler.FindAll)

	product := api.Group("/products", middleware.JWT())
	product.Post("/", productHandler.Create)
	product.Get("/", productHandler.FindAll)
	product.Get("/:id", productHandler.FindByID)
	product.Put("/:id", productHandler.Update)
	product.Delete("/:id", productHandler.Delete)

	address := api.Group("/addresses", middleware.JWT())
	address.Post("/", addressHandler.Create)
	address.Get("/", addressHandler.FindAll)
	address.Get("/:id", addressHandler.FindByID)
	address.Put("/:id", addressHandler.Update)
	address.Delete("/:id", addressHandler.Delete)

	cart := api.Group("/carts", middleware.JWT())
	cart.Post("/", cartHandler.AddToCart)
	cart.Get("/", cartHandler.FindAll)
	cart.Put("/:id", cartHandler.Update)
	cart.Delete("/:id", cartHandler.Delete)
}
