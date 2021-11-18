package main

import (
	"errors"
	"flag"

	"github.com/rs/zerolog/log"
)

var (
	configFile = flag.String("cfg", "", "YAML configuration file specifying Github repos to be backed up and backup target.")
)

func parseFlags() error {
	flag.Parse()

	log.Info().Msg("Config file name: " + *configFile)

	if stringIsNilOrEmpty(*configFile) {
		return errors.New("cfg cannot be empty")
	}

	return nil
}
