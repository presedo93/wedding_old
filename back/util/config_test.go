package util

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadEnvFile(t *testing.T) {
	tmp, err := os.CreateTemp(".", ".env")
	if err != nil {
		t.Fatal(err)
	}

	// Clean up
	defer os.Remove(tmp.Name())

	// Write some test data to it
	text := []byte("DATABASE_URL=test_db_url\nSERVER_ADDRESS=test_server_address")
	if _, err := tmp.Write(text); err != nil {
		t.Fatal(err)
	}

	conf, err := LoadEnv(".", tmp.Name())
	if err != nil {
		t.Fatal("Error loading .env file")
	}

	assert.Equal(t, "test_db_url", conf.DatabaseURL)
	assert.Equal(t, "test_server_address", conf.ServerAddress)
}
