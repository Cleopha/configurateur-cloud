package utils

import (
	"fmt"
	"github.com/rs/zerolog/log"
	"os"
)

var (
	ErrMissingEnvVar = func(name string) error {
		return fmt.Errorf("missing %s in environment", name)
	}
)

// MustGetEnv retrieves an environment variable value or exit program if the
// variable is missing.
func MustGetEnv(name string) string {
	value, exist := os.LookupEnv(name)
	if !exist {
		if os.Getenv("CI") == "" {
			log.Fatal().Err(ErrMissingEnvVar(name)).Msg("cannot start program")
		} else {
			log.Warn().Msg(ErrMissingEnvVar(name).Error())
		}
	}

	return value
}
