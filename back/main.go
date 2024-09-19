package main

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/presedo93/wedding/back/api"
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

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(conf.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
