package main

import (
	"context"
	"os"

	"github.com/MicahParks/keyfunc/v3"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/presedo93/wedding/back/api"
	"github.com/presedo93/wedding/back/auth"
	db "github.com/presedo93/wedding/back/db/sqlc"
	"github.com/presedo93/wedding/back/util"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	setupLogger()

	conf, err := util.LoadEnv(".")
	if err != nil {
		log.Fatal().Err(err).Msgf("cannot load env")
	}

	conn, err := pgxpool.New(context.Background(), conf.DatabaseURL)
	if err != nil {
		log.Fatal().Err(err).Msgf("cannot connect to db")
	}

	keyfunc, err := keyfunc.NewDefault([]string{conf.JwksURL})
	if err != nil {
		log.Fatal().Err(err).Msgf("cannot create jwks")
	}

	jwks := auth.NewJWKS(keyfunc, conf.IssuerURL)
	store := db.NewStore(conn)
	server := api.NewServer(store, jwks)

	err = server.Start(conf.ServerAddress)
	if err != nil {
		log.Fatal().Err(err).Msgf("cannot start server")
	}
}

func setupLogger() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout})
}
