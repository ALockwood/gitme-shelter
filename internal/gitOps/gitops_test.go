package gitOps

import (
	"strings"
	"testing"
)

func TestNewGitBundler(t *testing.T) {
	r := []GitRepo{{
		Name: "testRepo",
		URI:  "notarealuri",
	}}
	d := "workingDir"

	gb := NewGitBundler(&r, d)
	if gb.Bundler().workDir != d {
		t.Fatal("Failed to set workdir correctly")
	}

	if len(*gb.Bundler().GitRepos) != 1 {
		t.Fatal("Failed to set git repos")
	}

	//Kind of stupid because it'll fail if git isn't installed
	if strings.TrimSpace(gb.Bundler().gitPath) == "" {
		t.Error("Failed to locate or set gitpath")
	}
}
