package main

import (
	"errors"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Config struct {
	S3Bucket   string       `yaml:"s3Bucket"`
	AwsRegion  string       `yaml:"awsRegion"`
	GithubRepo []githubRepo `yaml:"githubRepo"`
}

func loadConfig() (*Config, error) {
	//Parse command line args
	err := parseFlags()
	if err != nil {
		return nil, err
	}

	var cfg Config
	cfgData, err := ioutil.ReadFile(*configFile)
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
