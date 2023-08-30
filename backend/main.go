package main

import (
	"backend/cmd"
	"backend/logger"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"os"
)

func main() {
	defaultLogger, err := logger.New()
	if err != nil {
		panic(err)
	}

	// Replace default logger with a custom one.
	logLevel := zerolog.DebugLevel
	if os.Getenv("GIN_MODE") == "release" {
		logLevel = zerolog.InfoLevel
	}

	log.Logger = defaultLogger.Level(logLevel)

	if err := cmd.Execute(); err != nil {
		log.Fatal().Err(err).Msg("cannot execute command")
	}
}
