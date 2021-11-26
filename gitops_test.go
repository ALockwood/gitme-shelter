package main

import (
	"strings"
	"testing"
)

func TestNewGitBundler(t *testing.T) {
	r := []gitRepo{{
		Name: "testRepo",
		URI:  "notarealuri",
	}}
	d := "workingDir"

	gb := newGitBundler(&r, d)
	if gb.bundler().workDir != d {
		t.Fatal("Failed to set workdir correctly")
	}

	if len(*gb.bundler().gitRepos) != 1 {
		t.Fatal("Failed to set git repos")
	}

	//Kind of stupid because it'll fail if git isn't installed
	if strings.TrimSpace(gb.bundler().gitPath) == "" {
		t.Error("Failed to locate or set gitpath")
	}
}
