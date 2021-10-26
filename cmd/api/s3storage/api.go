package s3storage

import (
	"github.com/pkg/errors"
)

type Storage interface {
	Upload(file S3File) (string, error)
	Read(url string) (S3File, error)
}

type Options struct {
	Bucket   S3Bucket
	S3Client S3ClientOptions
}

type storage struct {
	bucket   S3Bucket
	s3Client S3Client
}

func NewStorage(opt Options) (Storage, error) {
	s3Client, err := NewS3Client(opt.S3Client)
	if err != nil {
		return nil, errors.Wrap(err, "create_s3_storage")
	}

	return &storage{
		bucket:   opt.Bucket,
		s3Client: s3Client,
	}, nil
}

func (s *storage) Upload(file S3File) (string, error) {
	return s.s3Client.Upload(s.bucket, file)
}

func (s *storage) Read(url string) (S3File, error) {
	return s.s3Client.DownloadByUrl(s.bucket, url)
}
