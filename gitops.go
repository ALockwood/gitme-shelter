package main

import (
	"os"
	"os/exec"

	"github.com/rs/zerolog/log"
)

type gitBundler struct {
	gitPath string
	workDir string
	config  *Config
}

func newGitBundler(cfg *Config, workingDir string) *gitBundler {
	return &gitBundler{
		gitPath: getGitPath(),
		workDir: workingDir,
		config:  cfg}
}

func getGitPath() string {
	gitPath, err := exec.LookPath("git")
	if err != nil {
		log.Fatal().Msg("Failed to locat git. Is git installed?")
	}

	return gitPath
}

func (gb gitBundler) getBundles() error {
	if stringIsNilOrEmpty(gb.gitPath) || gb.config == nil || stringIsNilOrEmpty(gb.workDir) {
		log.Fatal().Msg("gitBundler not initialized")
	}

	for _, repo := range gb.config.GithubRepo {
		log.Debug().Msg(repo.Name)
		bcmd := exec.Cmd{
			Path:   gb.gitPath,
			Dir:    gb.workDir,
			Args:   []string{gb.gitPath, "version"},
			Stdout: os.Stdout,
			Stdin:  os.Stdin,
		}

		//log.Info().Msg(bcmd.String())
		if err := bcmd.Run(); err != nil {
			log.Fatal().Msg(err.Error())
		}
	}

	return nil
}
