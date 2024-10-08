package main

import (
	"context"
	"log"

	"github.com/MicahParks/keyfunc/v3"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/presedo93/wedding/back/api"
	"github.com/presedo93/wedding/back/auth"
	db "github.com/presedo93/wedding/back/db/sqlc"
	"github.com/presedo93/wedding/back/util"
)

func main() {
	conf, err := util.LoadEnv(".")
	if err != nil {
		log.Fatal("cannot load env:", err)
	}

	conn, err := pgxpool.New(context.Background(), conf.DatabaseURL)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	keyfunc, err := keyfunc.NewDefault([]string{conf.JwksURL})
	if err != nil {
		log.Fatal("cannot create jwks:", err)
	}

	jwks := auth.NewJWKS(keyfunc, conf.IssuerURL)
	store := db.NewStore(conn)
	server := api.NewServer(store, jwks)

	err = server.Start(conf.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
