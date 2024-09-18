package db

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/presedo93/wedding/back/util"
)

var testStore Store

func TestMain(m *testing.M) {
	conf, err := util.LoadEnv("../..", ".env")
	if err != nil {
		log.Fatal("cannot load env:", err)
	}

	connPool, err := pgxpool.New(context.Background(), conf.DatabaseURL)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	testStore = NewStore(connPool)
	os.Exit(m.Run())
}
