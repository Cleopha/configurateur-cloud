package common

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// GetRequestLogger returns a contextual logger with information related to
// the request itself.
//
// This is useful for monitoring and debug purpose.
func GetRequestLogger(c *gin.Context) zerolog.Logger {
	contextLogger, exist := c.Get("logger")
	if !exist {
		log.Debug().Msg("could not found contextual logger, use default logger")
		return log.Logger
	}

	logger, ok := contextLogger.(zerolog.Logger)
	if !ok {
		log.Debug().Msg("contextual logger is not valid, use default logger")
	}

	return logger
}
