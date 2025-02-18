package auth

import (
	"github.com/MicahParks/keyfunc/v3"
	"github.com/golang-jwt/jwt/v5"
)

type JWKS interface {
	VerifyToken(tokenString string) (jwt.MapClaims, error)
}

type LogtoJWKS struct {
	keyfunc keyfunc.Keyfunc
}

func (l *LogtoJWKS) VerifyToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, l.keyfunc.Keyfunc)
	if err != nil || !token.Valid {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, err
	}

	return claims, nil
}

func NewJWKS(keyfunc keyfunc.Keyfunc) JWKS {
	return &LogtoJWKS{
		keyfunc: keyfunc,
	}
}
