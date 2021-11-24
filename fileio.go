package main

import (
	"os"
	"path/filepath"

	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
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

//Remove the directory passed in
func removeWorkingDir(wrkdir string) error {
	log.Debug().Msgf("Removing %s and all child objects", wrkdir)
	if err := os.RemoveAll(wrkdir); err != nil {
		return err
	}
	log.Debug().Msg("All temp files deleted")
	return nil
}
