package main

import (
	"errors"
	"flag"
	"fmt"
	"strings"
)

var (
	cfg    = flag.String("cfg", "", "YAML configuration file specifying Github repos to be backed up and backup target.")
	dryrun = flag.Bool("dryrun", false, "Process config and check access for repo(s) and storage, but don't perform backup(s).")
)

func getFlags() error {
	flag.Parse()

	fmt.Println("Config file name: ", *cfg)
	fmt.Println("Dry run?: ", *dryrun)

	if strings.TrimSpace(*cfg) == "" {
		return errors.New("cfg cannot be empty")
	}

	return nil
}
