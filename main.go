package main

import (
	"flag"
	"fmt"
	"log"
)

var (
	cfg    = flag.String("cfg", "", "YAML configuration file specifying Github repos to be backed up and backup target.")
	dryrun = flag.Bool("dryrun", false, "Process config and check access for repo(s) and storage, but don't perform backup(s).")
)

func main() {
	//Check command line args
	err := parseFlags()
	if err != nil {
		log.Fatal(err)
	}

	//Load config
	conf, err := parseConfig("default-test.yaml")
	if err != nil {
		log.Fatal("failed to parse config")
	}
	fmt.Println(conf)

	//Validate S3 access
	//Main Loop
	//	Validate access to git repo
	//	check/create temp dir
	//	do git bundle
	//	compress
	//	upload to S3
}
