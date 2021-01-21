package main

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("hello")
	})

	app.Listen(fmt.Sprintf(":%s", os.Getenv("PORT")))
}
