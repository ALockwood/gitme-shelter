package main

import (
	"os"
	"path/filepath"

	"github.com/google/uuid"
)

//To avoid conflicts with other runs, generate a UUID and use it to create a temp dir
func createWorkingDir() (string, error) {
	guid := uuid.New()
	tpath := filepath.Clean(filepath.Join(os.TempDir(), "gitmeShelter-"+guid.String()))

	if err := os.Mkdir(tpath, os.ModePerm); err != nil {
		return "", err
	}

	return tpath, nil
}
