package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

// GenerateRequestID binds an identifier to the request so further operation
// that happens in this request can be tracked.
func GenerateRequestID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//	requestID := uuid.New()

		logger := log.With().
			//	Str("request_id", requestID.String()).
			Logger()

		ctx.Set("logger", logger)

		logger.Info().Msg("New request received")

		ctx.Next()
	}
}
