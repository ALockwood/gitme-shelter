package main

import (
	"errors"
	"flag"

	"github.com/rs/zerolog/log"
)

var (
	cfg    = flag.String("cfg", "", "YAML configuration file specifying Github repos to be backed up and backup target.")
	dryrun = flag.Bool("dryrun", false, "Process config and check access for repo(s) and storage, but don't perform backup(s).")
)

func parseFlags() error {
	flag.Parse()

	log.Info().Msg("Config file name: " + *cfg)
	log.Info().Msgf("Dry run?: %t", *dryrun)

	if stringIsNilOrEmpty(*cfg) {
		return errors.New("cfg cannot be empty")
	}

	return nil
}
