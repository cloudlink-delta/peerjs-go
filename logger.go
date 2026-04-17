package peer

import (
	"os"
	"time"

	"github.com/rs/zerolog"
)

func createLogger(source string, level zerolog.Level) zerolog.Logger {
	output := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}
	log := zerolog.New(output).With().Timestamp().Logger().Level(level)
	logger := log.With().
		Str("module", "peer").
		Str("source", source).
		Logger()
	return logger
}
