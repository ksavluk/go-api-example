package app

import (
	"ksavluk/go-api-example/cmd/api/cognito"
	"ksavluk/go-api-example/cmd/api/s3storage"
	"os"
)

type Options struct {
	UserManager     cognito.Options
	FileStorage     s3storage.Options
	ExampleHttpHost string
}

func GetEnvOptions() Options {
	cognitoOptions := cognito.Options{
		ClientID:  os.Getenv("COGNITO_CLIENT_ID"),
		PoolID:    os.Getenv("COGNITO_POOL_ID"),
		Region:    os.Getenv("COGNITO_REGION"),
		AccessKey: os.Getenv("AWS_ACCESS_KEY_ID"),
		SecretKey: os.Getenv("AWS_SECRET_ACCESS_KEY"),
	}

	s3StorageOptions := s3storage.Options{
		Bucket: s3storage.S3Bucket{
			Name: os.Getenv("S3_SOURCE_BUCKET_NAME"),
			Url:  os.Getenv("S3_SOURCE_BUCKET_URL"),
		},
		S3Client: s3storage.S3ClientOptions{
			Access:   os.Getenv("S3_ACCESS_KEY"),
			Secret:   os.Getenv("S3_SECRET_KEY"),
			Endpoint: os.Getenv("S3_ENDPOINT"),
			Region:   os.Getenv("S3_REGION"),
		},
	}

	return Options{
		UserManager:     cognitoOptions,
		FileStorage:     s3StorageOptions,
		ExampleHttpHost: os.Getenv("EXAMPLE_HTTP_HOST"),
	}
}
