package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/rs/zerolog/log"
)

type BundleUploader struct {
	S3Uploader *s3manager.Uploader
	GitBundler *gitBundler
}

func initUploader(gb *gitBundler) BundleUploader {
	sess := session.Must(session.NewSession(&aws.Config{Region: aws.String(gb.config.AwsRegion)}))
	s3Svc := s3.New(sess)

	u := s3manager.NewUploaderWithClient(s3Svc, func(u *s3manager.Uploader) {
		u.PartSize = 32 * 1024 * 1024 // 32MB per part
		u.LeavePartsOnError = false
		u.Concurrency = 2 //2 go routines per upload
	})

	return BundleUploader{
		S3Uploader: u,
		GitBundler: gb,
	}
}

func (bu *BundleUploader) UploadBundles() {
	var wg sync.WaitGroup
	t := time.Now().UTC()
	keyPfx := fmt.Sprintf("%d-%02d-%02d/%02d.%02d/", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute())

	for _, repo := range bu.GitBundler.config.GithubRepo {
		wg.Add(1)
		log.Trace().Msg(repo.TempDirectory)
		go bu.uploadBundle(filepath.Clean(filepath.Join(repo.TempDirectory, repo.BundleFile)), repo.BundleFile, keyPfx, &wg)
	}

	wg.Wait()
}

func (bu *BundleUploader) uploadBundle(bundle string, filename string, keyPrefix string, wg *sync.WaitGroup) {
	defer wg.Done()

	b, err := os.Open(bundle)
	if err != nil {
		log.Error().Msgf("failed to open file %q, %v", bundle, err)
	}

	log.Trace().Msg(bundle)
	log.Trace().Msg(filename)

	res, err := bu.S3Uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bu.GitBundler.config.S3Bucket),
		Key:    aws.String(keyPrefix + filename),
		Body:   b,
	})
	if err != nil {
		log.Error().Msgf("failed to upload file, %v", err)
	}

	log.Info().Msgf("Bundle uploaded to %s", res.Location)
}
