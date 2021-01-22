package main

import (
	"context"
	"fmt"
	"os"

	"github.com/go-pg/pg/v10"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var db *pg.DB

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	// TODO: Use a less powerful user and use a DB other than defaultdb.
	db = pg.Connect(&pg.Options{
		Addr:     "chalky-cockroachdb-public",
		User:     "root",
		Database: "defaultdb",
	})
	if err := db.Ping(context.Background()); err != nil {
		log.Fatal().Err(err).Msg("Failed to ping CockroachDB")
	}

	app := setupRouter()

	// TODO: Use TLS.
	addr := fmt.Sprintf(":%s", os.Getenv("PORT"))
	log.Fatal().Err(app.Listen(addr)).Str("addr", addr).Msg("Failed to listen")
}
