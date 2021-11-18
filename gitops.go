package main

import (
	"bytes"
	"os/exec"
	"path/filepath"

	"github.com/rs/zerolog/log"
)

type githubRepo struct {
	Name          string `yaml:"name"`
	URI           string `yaml:"uri"`
	TempDirectory string
	BundleFile    string
}

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

func (gb gitBundler) makeBundles() error {
	if stringIsNilOrEmpty(gb.gitPath) || gb.config == nil || stringIsNilOrEmpty(gb.workDir) {
		log.Fatal().Msg("gitBundler not initialized")
	}

	for _, repo := range gb.config.GithubRepo {
		if err := gb.cloneRepo(repo.Name, repo.URI); err != nil {
			log.Fatal().Msgf("Failed to clone repo %s", repo.Name)
		}

		if err := gb.bundleRepo(&repo); err != nil {
			log.Fatal().Msgf("Failed to bundle repo %s", repo.Name)
		}
	}

	return nil
}

func (gb gitBundler) cloneRepo(name string, uri string) error {
	log.Debug().Msgf("Cloning repo %s", name)

	var out bytes.Buffer
	var errOut bytes.Buffer

	bcmd := exec.Cmd{
		Path:   gb.gitPath,
		Dir:    gb.workDir,
		Args:   []string{gb.gitPath, "clone", uri, name},
		Stdout: &out,
		Stderr: &errOut,
	}

	log.Trace().Msg(bcmd.String())
	if err := bcmd.Run(); err != nil {
		log.Error().Msg(errOut.String())
		return err
	}

	return nil
}

func (gb gitBundler) bundleRepo(r *githubRepo) error {
	log.Debug().Msgf("Bundling repo %s", r.Name)

	var out bytes.Buffer
	var errOut bytes.Buffer
	r.TempDirectory = filepath.Clean(filepath.Join(gb.workDir, r.Name))
	r.BundleFile = r.Name + ".bundle"

	log.Trace().Msg(r.TempDirectory)
	bcmd := exec.Cmd{
		Path:   gb.gitPath,
		Dir:    r.TempDirectory,
		Args:   []string{gb.gitPath, "bundle", "create", r.BundleFile, "--all"},
		Stdout: &out,
		Stderr: &errOut,
	}

	log.Trace().Msg(bcmd.String())
	if err := bcmd.Run(); err != nil {
		log.Error().Msg(errOut.String())
		return err
	}
	log.Info().Msgf("Completed bundling repo %s", r.Name)

	out.Reset()
	errOut.Reset()

	log.Debug().Msgf("Verifying bundle %s", r.BundleFile)
	vcmd := exec.Cmd{
		Path:   gb.gitPath,
		Dir:    r.TempDirectory,
		Args:   []string{gb.gitPath, "bundle", "verify", r.BundleFile},
		Stdout: &out,
		Stderr: &errOut,
	}

	log.Trace().Msg(vcmd.String())
	if err := vcmd.Run(); err != nil {
		log.Error().Msg("Bundle verification failed!")
		log.Error().Msg(errOut.String())
		return err
	}
	log.Info().Msgf("Completed verifying repo %s", r.Name)

	return nil
}
