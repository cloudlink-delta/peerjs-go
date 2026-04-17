package server

import (
	"os"
	"time"

	"github.com/rs/zerolog"
)

// return time in millis
// credits https://stackoverflow.com/questions/24122821/go-golang-time-now-unixnano-convert-to-milliseconds
func getTime() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

func createLogger(ctx string, opts Options) zerolog.Logger {
	output := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}
	log := zerolog.New(output).With().Timestamp().Logger().Level(opts.LogLevel)

	logger := log.With().
		Str("module", "server").
		Str("context", ctx).
		Logger()

	return logger
}
