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

func main() {
	initLogger()
	//Parse command args + env vars, load config data
	conf, err := loadConfig()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to parse config.")
	}
	log.Debug().Msgf("%+v", *conf)

	//Test write access to temp working dir by creating it
	tmpdir, err := createWorkingDir()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to create temp working dir. Err:")
	}
	log.Info().Msgf("Running with working dir %s", tmpdir)

	gb := newGitBundler(&conf.GithubRepo, tmpdir)
	gb.makeBundles()

	//	log.Debug().Msgf("%+v", gb.config.GithubRepo)

	//	upload to S3
	u := newUploader(gb.bundler(), conf)
	u.UploadBundles()

	//remove working dir
	removeWorkingDir(tmpdir)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to remove working dir!")
	}
}
