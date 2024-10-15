package logs

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func Middleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		start := time.Now()
		ctx.Next()
		end := time.Now()
		latency := end.Sub(start)

		log.Info().
			Str("method", ctx.Request.Method).
			Str("path", ctx.Request.RequestURI).
			Int("status", ctx.Writer.Status()).
			Str("ip", ctx.ClientIP()).
			Dur("latency", latency).
			Msg("Request handled")
	}
}
