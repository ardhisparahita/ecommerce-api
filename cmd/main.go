package main

import (
	"log"

	"github.com/ardhisparahita/ecommerce-api/internal/handler"
	"github.com/ardhisparahita/ecommerce-api/internal/repository"
	"github.com/ardhisparahita/ecommerce-api/internal/routes"
	"github.com/ardhisparahita/ecommerce-api/internal/service"
	"github.com/ardhisparahita/ecommerce-api/pkg/config"
	"github.com/ardhisparahita/ecommerce-api/pkg/database"
	"github.com/gofiber/fiber/v2"
)

func main() {
	config.LoadEnv()

	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}

	_ = db

	app := fiber.New()

	userRepo := repository.NewUserRepository(db)
	userService := service.NewAuthService(userRepo)
	userHandler := handler.NewAuthHandler(userService)

	categoryRepo := repository.NewCategoryRepository(db)
	categoryService := service.NewCategoryService(categoryRepo)
	categoryHandler := handler.NewCategoryHandler(categoryService)

	routes.SetupRoutes(app, userHandler, categoryHandler)

	log.Fatal(
		app.Listen(":" + config.Get("APP_PORT")),
	)
}
