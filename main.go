package main

import (
	"fmt"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	app := setupRouter()

	// TODO: Use TLS.
	addr := fmt.Sprintf(":%s", os.Getenv("PORT"))
	log.Fatal().Err(app.Listen(addr)).Str("addr", addr).Msg("Failed to listen")
}
