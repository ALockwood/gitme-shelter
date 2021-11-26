package main

import (
	"bytes"
	"os/exec"
	"path/filepath"

	"github.com/rs/zerolog/log"
)

type gitRepo struct {
	Name          string `yaml:"name"`
	URI           string `yaml:"uri"`
	TempDirectory string
	BundleFile    string
}

type gitBundler struct {
	gitPath  string
	workDir  string
	gitRepos *[]gitRepo
}

type GithubRepoBundler interface {
	makeBundles() error
	bundler() *gitBundler
}

func newGitBundler(gitRepo *[]gitRepo, workingDir string) GithubRepoBundler {
	gitPath, err := exec.LookPath("git")
	if err != nil {
		log.Fatal().Msg("Failed to locate git. Is git installed?")
	}

	return &gitBundler{
		gitPath:  gitPath,
		workDir:  workingDir,
		gitRepos: gitRepo}
}

func (gb *gitBundler) makeBundles() error {
	if stringIsNilOrEmpty(gb.gitPath) || gb.gitRepos == nil || stringIsNilOrEmpty(gb.workDir) {
		log.Fatal().Msg("gitBundler not initialized")
	}

	//Using range here felt odd as it's copying the values which I would than assign back.
	//ex:
	// i , r + range *gb.gitRepos {
	// ...
	// (*gb.gitRepos)[i] = r
	// }
	//I suspect what I'm doing is stupid. I just don't know how to do correctly/idiomatically
	for i := range *gb.gitRepos {
		r := &(*gb.gitRepos)[i]
		if err := r.cloneRepo(gb); err != nil {
			log.Fatal().Msgf("Failed to clone repo %s", r.Name)
		}

		if err := r.bundleRepo(gb); err != nil {
			log.Fatal().Msgf("Failed to bundle repo %s", r.Name)
		}
	}

	return nil
}

func (gb *gitBundler) bundler() *gitBundler {
	return gb
}

func (r *gitRepo) cloneRepo(gb *gitBundler) error {
	log.Debug().Msgf("Cloning repo %s", r.Name)

	var out bytes.Buffer
	var errOut bytes.Buffer

	bcmd := exec.Cmd{
		Path:   gb.gitPath,
		Dir:    gb.workDir,
		Args:   []string{gb.gitPath, "clone", r.URI, r.Name},
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

func (r *gitRepo) bundleRepo(gb *gitBundler) error {
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
