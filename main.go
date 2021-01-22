package main

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var db *pgx.Conn
var infiniteContext = context.Background()

func main() {
	var err error
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	hostname, _ := os.Hostname()
	log.Info().Msg("I am " + hostname)

	// TODO: Use a DB other than defaultdb and a less powerful user.
	dbURL := fmt.Sprintf("postgres://root@chalky-cockroachdb-public:26257/defaultdb?sslmode=disable")
	log.Info().Str("url", dbURL).Msg("Connecting to CockroachDB")
	db, err = pgx.Connect(infiniteContext, dbURL)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to connect to CockroachDB")
	}

	app := setupRouter()

	// TODO: Use TLS.
	addr := fmt.Sprintf(":%s", os.Getenv("PORT"))
	log.Fatal().Err(app.Listen(addr)).Str("addr", addr).Msg("Failed to listen")
}
