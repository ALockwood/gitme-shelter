package main

import (
	"flag"
	"os"
	"testing"
)

func TestParseFlags(t *testing.T) {
	//Save and restore OG arguments
	oga := os.Args
	defer func() { os.Args = oga }()

	c := "testfile.cfg"
	os.Args = []string{"", "-cfg", c} //first arg is path to executable

	x, err := parseFlags()
	if err != nil {
		t.Fatal("Error encountered during command arg parsing")
	}

	if *x.configFile != c {
		t.Error("Failed to set config file via command args")
	}

	t.Cleanup(func() {
		flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError) //Reset flags
	})
}

func TestLoadConfig(t *testing.T) {
	//Save and restore OG arguments
	oga := os.Args
	defer func() { os.Args = oga }()

	c := "default-test.yaml"
	os.Args = []string{"", "-cfg", c} //first arg is path to executable

	_, err := loadConfig()

	if err != nil {
		t.Fatalf("Failed to load and/or parse test config file %s", c)
	}

	t.Cleanup(func() {
		flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError) //Reset flags
	})
}
