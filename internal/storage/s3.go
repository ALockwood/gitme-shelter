package storage

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"

	cfg "github.com/ALockwood/gitme-shelter/internal/configuring"
	g "github.com/ALockwood/gitme-shelter/internal/gitOps"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/rs/zerolog/log"
)

type S3BundleUploader struct {
	S3Uploader *s3manager.Uploader
	GitBundler *g.GitBundler
	S3Bucket   string
}

type BundleUploader interface {
	UploadBundles()
}

func NewUploader(gb *g.GitBundler, c *cfg.Config) BundleUploader {
	sess := session.Must(session.NewSession(&aws.Config{Region: aws.String(c.AwsRegion)}))
	s3Svc := s3.New(sess)

	u := s3manager.NewUploaderWithClient(s3Svc, func(u *s3manager.Uploader) {
		//u.PartSize = 128 * 1024 * 1024 // 32MB per part
		u.LeavePartsOnError = false
		//u.Concurrency = 8 //8 go routines per upload
	})

	return &S3BundleUploader{
		S3Uploader: u,
		GitBundler: gb,
		S3Bucket:   c.S3Bucket,
	}
}

func (bu *S3BundleUploader) UploadBundles() {
	var wg sync.WaitGroup
	t := time.Now().UTC()
	keyPfx := fmt.Sprintf("%d-%02d-%02d/%02d.%02d/", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute())

	for _, repo := range *bu.GitBundler.GitRepos {
		wg.Add(1)
		log.Trace().Msg(repo.TempDirectory)
		go bu.uploadBundle(filepath.Clean(filepath.Join(repo.TempDirectory, repo.BundleFile)), repo.BundleFile, keyPfx, &wg)
	}

	wg.Wait()
}

func (bu *S3BundleUploader) uploadBundle(bundle string, filename string, keyPrefix string, wg *sync.WaitGroup) {
	defer wg.Done()

	b, err := os.Open(bundle)
	if err != nil {
		log.Error().Msgf("failed to open file %q, %v", bundle, err)
	}

	log.Info().Msgf("Uploading %s", bundle)

	res, err := bu.S3Uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bu.S3Bucket),
		Key:    aws.String(keyPrefix + filename),
		Body:   b,
	})
	if err != nil {
		log.Error().Msgf("failed to upload file, %v", err)
	}

	log.Info().Msgf("Bundle uploaded to %s", res.Location)
}
