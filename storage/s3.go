package storage

import (
	"bytes"
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

// S3 Helper
type s3helper struct {
	client *s3.Client
}

func newS3Helper() (s3helper, error) {
	cfg, err := config.LoadDefaultConfig(context.Background())
	if err != nil {
		return s3helper{}, err
	}
	s3Client := s3.NewFromConfig(cfg)
	return s3helper{
		client: s3Client,
	}, nil
}

// S3 storage
type S3 struct {
	*storageBase
	helper s3helper
	bucket string
}

func (s *S3) Upload(content []byte, id string) error {
	// TODO
	// Upload files to S3
	uploader := manager.NewUploader(s.helper.client)
	uploader.Upload(context.Background(), &s3.PutObjectInput{
		Bucket: aws.String(s.bucket),
		Key:    aws.String(id),
		Body:   bytes.NewReader(content),
	})
	return nil
}

func (s *S3) Download(id string) ([]byte, error) {
	downloader := manager.NewDownloader(s.helper.client)
	var buf []byte
	buff := manager.NewWriteAtBuffer(buf)
	_, err := downloader.Download(context.Background(), buff, &s3.GetObjectInput{
		Bucket: aws.String(s.bucket),
		Key:    aws.String(id),
	})
	if err != nil {
		return []byte(""), nil
	}
	return buff.Bytes(), nil
}

func NewS3() (S3, error) {
	s3h, err := newS3Helper()
	if err != nil {
		return S3{}, err
	}
	return S3{
		helper: s3h,
	}, nil
}
