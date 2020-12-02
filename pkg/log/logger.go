package log

import (
	"github.com/rs/zerolog"
	zlog "github.com/rs/zerolog/log"
	"os"
)

var logger zerolog.Logger

func init() {
	logger = zlog.Output(zerolog.ConsoleWriter{
		Out:        os.Stderr,
		TimeFormat: "15:04:05",
	})
}

func Logger() zerolog.Logger {
	return logger
}
