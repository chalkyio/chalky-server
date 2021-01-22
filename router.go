package main

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func setupRouter() *fiber.App {
	app := fiber.New()
	app.Use(logger.New())

	api := app.Group("/api")
	{
		api.Get("/health", func(c *fiber.Ctx) error {
			// We're up!
			return c.SendStatus(http.StatusOK)
		})
	}

	return app
}
