package logger

import (
	"fmt"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
	"path/filepath"
)

const (
	DefaultLogFile = "cloud-calculator-backend.log"
)

var (
	LogFolder = "logs"
)

func New() (zerolog.Logger, error) {
	file, err := os.OpenFile(filepath.Join(LogFolder, DefaultLogFile), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0664)
	if err != nil {
		return zerolog.Logger{}, err
	}

	consoleWriter := zerolog.ConsoleWriter{
		Out: os.Stdout,
	}

	multi := zerolog.MultiLevelWriter(consoleWriter, file)

	return zerolog.New(multi).With().
		Timestamp().
		Logger().
		Level(zerolog.InfoLevel), nil
}

func init() {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}

	LogFolder = filepath.Join(filepath.Dir(ex), LogFolder)

	log.Info().Msg(fmt.Sprintf("Log is save in %s", LogFolder))

	if err = os.MkdirAll(LogFolder, 0750); err != nil {
		panic(err)
	}
}
