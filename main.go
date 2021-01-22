package main

import (
	"context"
	"fmt"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/cockroachdb"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v4"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var db *pgx.Conn
var infiniteContext = context.Background()

const databaseURI = "root@chalky-cockroachdb-public:26257"

func main() {
	var err error
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	hostname, _ := os.Hostname()
	log.Info().Msg("I am " + hostname)

	log.Info().Msg("Attempting to migrate database")
	migrator, err := migrate.New(
		"file://migrations",
		"cockroachdb://"+databaseURI+"/defaultdb?sslmode=disable",
	)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to create migration manager")
	}
	if err := migrator.Up(); err != nil {
		log.Fatal().Err(err).Msg("Failed to migrate database")
	}

	log.Info().Str("uri", databaseURI).Msg("Connecting to CockroachDB")
	db, err = pgx.Connect(infiniteContext, "postgres://"+databaseURI+"/chalky?sslmode=disable")
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to connect to CockroachDB")
	}

	app := setupRouter()

	// TODO: Use TLS.
	addr := fmt.Sprintf(":%s", os.Getenv("PORT"))
	log.Fatal().Err(app.Listen(addr)).Str("addr", addr).Msg("Failed to start listening")
}
