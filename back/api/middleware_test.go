package api

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/presedo93/wedding/back/auth"
)

type MockJWKS struct {
	claims jwt.MapClaims
}

func (m *MockJWKS) VerifyToken(tokenString string) (jwt.MapClaims, error) {
	return m.claims, nil
}

func NewMockJWKS(user string) auth.JWKS {
	return &MockJWKS{claims: jwt.MapClaims{"sub": user}}
}
