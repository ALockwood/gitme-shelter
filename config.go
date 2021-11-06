package main

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type Config struct {
	S3Bucket   string `yaml:"s3Bucket"`
	GithubRepo []struct {
		Name string `yaml:"name"`
		URI  string `yaml:"uri"`
	} `yaml:"githubRepo"`
}

func parseConfig(cfgFile string) (*Config, error) {
	var cfg Config
	cfgData, err := ioutil.ReadFile(cfgFile)
	if err != nil {
		log.Fatal("failed to locate and/or read config file")
	}

	err = yaml.Unmarshal(cfgData, &cfg)
	if err != nil {
		log.Fatal("failed to parse config file")
	}

	return &cfg, nil
}
