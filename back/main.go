package main

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/presedo93/wedding/back/api"
	db "github.com/presedo93/wedding/back/db/sqlc"
)

const (
	dbSource      = "postgresql://rendres:s3cr3t@localhost:5432/wedding?sslmode=disable"
	serverAddress = "0.0.0.0:8080"
)

func main() {
	conn, err := pgxpool.New(context.Background(), dbSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(serverAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
