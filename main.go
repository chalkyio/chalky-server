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

	// TODO: Use a DB other than defaultdb and a less powerful user.
	dbURL := fmt.Sprintf("postgres://root@chalky-cockroachdb-public/defaultdb")
	db, err = pgx.Connect(context.Background(), dbURL)
	if err != nil {
		log.Fatal().Err(err).Str("url", dbURL).Msg("Failed to connect to CockroachDB")
	}

	app := setupRouter()

	// TODO: Use TLS.
	addr := fmt.Sprintf(":%s", os.Getenv("PORT"))
	log.Fatal().Err(app.Listen(addr)).Str("addr", addr).Msg("Failed to listen")
}
