package main

import (
	"log"

	"github.com/ardhisparahita/ecommerce-api/internal/handler"
	"github.com/ardhisparahita/ecommerce-api/internal/repository"
	"github.com/ardhisparahita/ecommerce-api/internal/routes"
	"github.com/ardhisparahita/ecommerce-api/internal/service"
	"github.com/ardhisparahita/ecommerce-api/pkg/config"
	"github.com/ardhisparahita/ecommerce-api/pkg/database"
	"github.com/ardhisparahita/ecommerce-api/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

func main() {
	config.LoadEnv()

	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}

	_ = db

	app := fiber.New(
		fiber.Config{
			ErrorHandler: utils.ErrorHandler,
		},
	)

	userRepo := repository.NewUserRepository(db)
	userService := service.NewAuthService(userRepo)
	userHandler := handler.NewAuthHandler(userService)

	categoryRepo := repository.NewCategoryRepository(db)
	categoryService := service.NewCategoryService(categoryRepo)
	categoryHandler := handler.NewCategoryHandler(categoryService)

	productRepo := repository.NewProductRepository(db)
	productService := service.NewProductService(productRepo)
	productHandler := handler.NewProductHandler(productService)

	addressRepo := repository.NewAddressRepository(db)
	addressService := service.NewAddressService(addressRepo)
	addressHandler := handler.NewAddressHandler(addressService)

	cartRepo := repository.NewCartRepository(db)
	cartService := service.NewCartService(cartRepo)
	cartHandler := handler.NewCartHandler(cartService)

	orderRepo := repository.NewOrderRepository(db)
	orderItemRepo := repository.NewOrderItemRepository(db)
	paymentRepo := repository.NewPaymentRepository(db)

	checkoutService := service.NewCheckoutService(db, cartRepo, productRepo, addressRepo, orderRepo, orderItemRepo, paymentRepo)
	checkoutHandler := handler.NewCheckoutHandler(checkoutService)

	orderService := service.NewOrderService(db, orderRepo, productRepo, paymentRepo)
	orderHandler := handler.NewOrderHandler(orderService)

	routes.SetupRoutes(app, userHandler, categoryHandler, productHandler, addressHandler, cartHandler, checkoutHandler, orderHandler)

	log.Fatal(
		app.Listen(":" + config.Get("APP_PORT")),
	)
}
