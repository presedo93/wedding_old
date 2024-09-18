package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadEnvVars(t *testing.T) {
	conf, err := LoadEnv("../", ".env")
	if err != nil {
		t.Fatal("Error loading .env file")
	}

	assert.Equal(t, "postgresql://rendres:s3cr3t@localhost:5432/wedding?sslmode=disable", conf.DatabaseURL)
	assert.Equal(t, "0.0.0.0:8080", conf.ServerAddress)
}
