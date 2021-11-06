package main

import (
	"errors"
	"flag"
	"fmt"
	"strings"
)

func parseFlags() error {
	flag.Parse()

	fmt.Println("Config file name: ", *cfg)
	fmt.Println("Dry run?: ", *dryrun)

	if strings.TrimSpace(*cfg) == "" {
		return errors.New("cfg cannot be empty")
	}

	return nil
}
