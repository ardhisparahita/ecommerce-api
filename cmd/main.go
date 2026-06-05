package main

import (
	"log"

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

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "e-commerce api running",
		})
	})

	log.Fatal(
		app.Listen(":" + config.Get("APP_PORT")),
	)
}
