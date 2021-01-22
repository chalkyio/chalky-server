package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	app := fiber.New()

	api := app.Group("/api")
	{
		api.Get("/health", func(c *fiber.Ctx) error {
			return c.SendStatus(http.StatusOK)
		})
	}

	// TODO: Use TLS.
	addr := fmt.Sprintf(":%s", os.Getenv("PORT"))
	log.Fatal().Err(app.Listen(addr)).Str("addr", addr).Msg("Failed to listen")
}
