package main

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

type statusCode int

const (
	errNotExists statusCode = iota
	errIncorrectPassword
	okAuthenticated
)

type responseBody struct {
	Code    statusCode  `json:"code"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func errorHandler(ctx *fiber.Ctx, err error) error {
	code := http.StatusInternalServerError
	if e, ok := err.(*fiber.Error); ok {
		// Override the status code.
		code = e.Code
	}

	log.Error().Err(err).Msg("Error returned from handler function")
	return ctx.SendStatus(code)
}
