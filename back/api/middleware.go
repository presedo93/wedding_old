package api

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/presedo93/wedding/back/auth"
)

const (
	authHeader = "Authorization"
	authBearer = "Bearer"
)

// AuthMiddleware creates a gin middleware for authorization
func authMiddleware(jwks auth.JWKS) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString, err := extractBearer(ctx.Request.Header)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(err))
			return
		}

		claims, err := jwks.VerifyToken(tokenString)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(err))
			return
		}

		// Get user ID from token
		userId := claims["sub"].(string)
		ctx.Set("userID", userId)

		// ctx.Set(authorizationPayloadKey, token)
		ctx.Next()
	}
}

func extractBearer(headers http.Header) (string, error) {
	authHeader := headers.Get(authHeader)
	if strings.HasPrefix(authHeader, authBearer) {
		return strings.TrimPrefix(authHeader, "Bearer "), nil
	}

	return "", errors.New("authorization header is incorrect")
}
