package s3storage

import (
	"bytes"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/pkg/errors"
	"io/ioutil"
	"strings"
)

type S3Bucket struct {
	Name, Url string
}

type S3File struct {
	Content           []byte
	Name, ContentType string
}

type S3Client interface {
	DownloadByUrl(bucket S3Bucket, url string) (S3File, error)
	Download(bucket S3Bucket, fileName string) (S3File, error)
	Upload(bucket S3Bucket, file S3File) (string, error)
}

type s3Client struct {
	*s3.S3
}

type S3ClientOptions struct {
	Access, Secret, Endpoint, Region string
}

func NewS3Client(opt S3ClientOptions) (S3Client, error) {
	enableSSL := len(opt.Endpoint) == 0

	s3Config := &aws.Config{
		Credentials:      credentials.NewStaticCredentials(opt.Access, opt.Secret, ""),
		Endpoint:         aws.String(opt.Endpoint),
		Region:           aws.String(opt.Region),
		DisableSSL:       aws.Bool(!enableSSL),
		S3ForcePathStyle: aws.Bool(true),
	}
	newSession, err := session.NewSession(s3Config)
	if err != nil {
		return nil, errors.Wrap(err, "create_s3_client")
	}

	return &s3Client{
		s3.New(newSession),
	}, nil
}

func (s *s3Client) Download(bucket S3Bucket, fileName string) (S3File, error) {
	output, err := s.S3.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(bucket.Name),
		Key:    aws.String(fileName),
	})
	if err != nil {
		return S3File{}, err
	}

	return s.outputToFile(fileName, output)
}

func (s *s3Client) DownloadByUrl(bucket S3Bucket, url string) (S3File, error) {
	return s.Download(bucket, strings.TrimPrefix(url, bucket.Url+"/"))
}

func (s *s3Client) Upload(bucket S3Bucket, file S3File) (string, error) {
	_, err := s.S3.PutObject(&s3.PutObjectInput{
		Bucket:        aws.String(bucket.Name),
		Key:           aws.String(file.Name),
		Body:          bytes.NewReader(file.Content),
		ContentLength: aws.Int64(int64(len(file.Content))),
		ContentType:   aws.String(file.ContentType),
	})
	if err != nil {
		return "", err
	}

	url := bucket.Url + "/" + file.Name
	return url, nil
}

func (s *s3Client) outputToFile(fileName string, o *s3.GetObjectOutput) (S3File, error) {
	body, err := ioutil.ReadAll(o.Body)
	if err != nil {
		return S3File{}, err
	}

	contentType := ""
	if o.ContentType != nil {
		contentType = *o.ContentType
	}

	return S3File{
		Content:     body,
		Name:        fileName,
		ContentType: contentType,
	}, nil
}
