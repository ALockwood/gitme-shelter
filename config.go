package main

import (
	"errors"
	"flag"
	"io/ioutil"

	"github.com/rs/zerolog/log"
	"gopkg.in/yaml.v2"
)

type CommandArgs struct {
	configFile *string
}

type Config struct {
	S3Bucket   string    `yaml:"s3Bucket"`
	AwsRegion  string    `yaml:"awsRegion"`
	GithubRepo []gitRepo `yaml:"githubRepo"`
}

func loadConfig() (*Config, error) {
	//Parse command line args
	args, err := parseFlags()
	if err != nil {
		return nil, err
	}

	var cfg Config
	cfgData, err := ioutil.ReadFile(*args.configFile)
	if err != nil {
		return nil, errors.New("failed to locate and/or read config file")
	}

	err = yaml.Unmarshal(cfgData, &cfg)
	if err != nil {
		return nil, err
	}

	if stringIsNilOrEmpty(cfg.S3Bucket) {
		return nil, errors.New("failed to load S3 bucket")
	}

	if len(cfg.GithubRepo) == 0 {
		return nil, errors.New("failed to load Github repo(s)")
	}

	return &cfg, nil
}

func parseFlags() (*CommandArgs, error) {
	cmdArgs := CommandArgs{
		configFile: flag.String("cfg", "", "YAML configuration file specifying Github repos to be backed up and backup target."),
	}
	flag.Parse()

	if stringIsNilOrEmpty(*cmdArgs.configFile) {
		return nil, errors.New("cfg cannot be empty")
	}
	log.Info().Msgf("Config file name: %s", *cmdArgs.configFile)

	return &cmdArgs, nil
}
