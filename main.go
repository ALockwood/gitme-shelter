package main

import (
	"github.com/rs/zerolog/log"
)

func main() {
	initLogger()
	//Parse command args + env vars, load config data
	conf, err := loadConfig("default-test.yaml")
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to parse config.")
	}
	log.Info().Msgf("%+v", *conf)

	//Test write access to temp working dir by creating it
	tmpdir, err := createWorkingDir()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to create temp working dir. Err:")
	}
	log.Info().Msg("Working dir: " + tmpdir)

	//Validate S3 access
	//Main Loop
	//	Validate access to git repo
	//	check/create temp dir
	//	do git bundle
	//	compress
	//	upload to S3
}
