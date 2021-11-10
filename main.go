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
	log.Debug().Msgf("%+v", *conf)

	//Test write access to temp working dir by creating it
	tmpdir, err := createWorkingDir()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to create temp working dir. Err:")
	}
	log.Info().Msgf("Running with working dir %s", tmpdir)

	//Validate S3 access
	//Main Loop
	//	Validate access to git repo
	//	check/create temp dir
	//	do git bundle
	//	compress
	//	upload to S3

	//remove working dir
	removeWorkingDir(tmpdir)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to remove working dir!")
	}
}
