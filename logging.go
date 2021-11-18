package main

import (
	"os"
	"strings"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func initLogger() {
	loglvl := strings.ToLower(getEnv("LOG_LEVEL", "debug"))
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Logger = log.Output((zerolog.ConsoleWriter{Out: os.Stderr}))

	switch loglvl {
	case "trace":
		zerolog.SetGlobalLevel(zerolog.TraceLevel)
	case "debug":
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	case "info":
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	case "warn":
		zerolog.SetGlobalLevel(zerolog.WarnLevel)
	case "error":
		zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	default:
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	}

	log.Info().Msgf("Log level set to %s", zerolog.GlobalLevel().String())
}
