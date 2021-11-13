package main

import (
	"errors"
	"io/ioutil"

	"github.com/rs/zerolog/log"
	"gopkg.in/yaml.v2"
)

type Config struct {
	S3Bucket   string `yaml:"s3Bucket"`
	GithubRepo []struct {
		Name string `yaml:"name"`
		URI  string `yaml:"uri"`
	} `yaml:"githubRepo"`
}

type Credentials struct {
	GithubToken        string
	AwsAccessKey       string
	AwsSecretAccessKey string
}

const envGithub string = "GITHUB_ACCESS_TOKEN"
const envAwsAccessKey string = "AWS_ACCESS_KEY_ID"
const envAwsSecretAccessKey = "AWS_SECRET_ACCESS_KEY"

func loadConfig(cfgFile string) (*Config, error) {
	//Parse command line args
	err := parseFlags()
	if err != nil {
		return nil, err
	}

	var cfg Config
	cfgData, err := ioutil.ReadFile(cfgFile)
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

	cred := Credentials{
		GithubToken:        getEnv(envGithub, ""),
		AwsAccessKey:       getEnv(envAwsAccessKey, ""),
		AwsSecretAccessKey: getEnv(envAwsSecretAccessKey, "")}
	log.Debug().Msgf("%+v", cred)
	/*
		if stringIsNilOrEmpty(cred.GithubToken) {
			return nil, errors.New(envGithub + " env var empty or unset")
		}
		if stringIsNilOrEmpty(cred.AwsAccessKey) {
			return nil, errors.New(envAwsAccessKey + " env var empty or unset")
		}
		if stringIsNilOrEmpty(cred.AwsSecretAccessKey) {
			return nil, errors.New(envAwsSecretAccessKey + " env var empty or unset")
		}
	*/
	return &cfg, nil
}
