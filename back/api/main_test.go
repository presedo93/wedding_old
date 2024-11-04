package api

import (
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/presedo93/wedding/back/auth"
	"github.com/rs/zerolog"
)

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	zerolog.SetGlobalLevel(zerolog.Disabled)
	os.Exit(m.Run())
}

// MockJWKS is a mock implementation of the JWKS interface.
type MockJWKS struct {
	claims jwt.MapClaims
}

func (m *MockJWKS) VerifyToken(tokenString string) (jwt.MapClaims, error) {
	return m.claims, nil
}

func NewMockJWKS(user uuid.UUID) auth.JWKS {
	return &MockJWKS{claims: jwt.MapClaims{"sub": user}}
}
