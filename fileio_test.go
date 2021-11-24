package main

import (
	"errors"
	"os"
	"strings"
	"testing"
)

func TestCreateWorkingDir(t *testing.T) {
	s, err := createWorkingDir()

	if strings.TrimSpace(s) == "" || err != nil {
		t.Errorf("Failed to create workdir")
	}

	os.RemoveAll(s)
}

func TestRemoveWorkingDir(t *testing.T) {
	d, _ := createWorkingDir()

	err := removeWorkingDir(d)
	if err != nil {
		t.Fatalf("Failed to delete temp working dir %s", d)
	}

	_, err = os.Stat(d)
	if err != nil && errors.Is(err, os.ErrNotExist) {
		return
	} else {
		t.Error("Failed to remote temp working dir")
	}
}
