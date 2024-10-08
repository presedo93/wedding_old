package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadEnvFile(t *testing.T) {
	conf, err := LoadEnv("../")
	if err != nil {
		t.Fatal("Error loading .env file")
	}

	assert.Equal(t, "postgresql://rendres:s3cr3t@localhost:5432/wedding?sslmode=disable", conf.DatabaseURL)
	assert.Equal(t, "0.0.0.0:8080", conf.ServerAddress)

	assert.Equal(t, "http://localhost:3001/oidc/jwks", conf.JwksURL)
	assert.Equal(t, "http://localhost:3001/oidc", conf.IssuerURL)
}
