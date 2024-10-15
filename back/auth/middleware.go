package auth

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

const (
	Header = "Authorization"
	Bearer = "Bearer"
)

// AuthMiddleware creates a gin middleware for authorization
func Middleware(jwks JWKS) gin.HandlerFunc {
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
	authHeader := headers.Get(Header)
	if strings.HasPrefix(authHeader, Bearer) {
		return strings.TrimPrefix(authHeader, "Bearer "), nil
	}

	return "", errors.New("authorization header is incorrect")
}

func errorResponse(err error) gin.H {
	log.Error().Err(err).Msg("JWT Auth")
	return gin.H{"error": err.Error()}
}
