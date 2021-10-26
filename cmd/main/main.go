package main

import (
	"ksavluk/go-api-example/cmd/api/cognito"
	"ksavluk/go-api-example/cmd/api/httpExample"
	"ksavluk/go-api-example/cmd/api/s3storage"
	"ksavluk/go-api-example/cmd/app"
	"log"
)

func main() {
	log.Println("Read environment...")
	opt := app.GetEnvOptions()

	tryHttpApi(opt.ExampleHttpHost)
	tryCognitoUserApi(opt.UserManager)
	tryS3Storage(opt.FileStorage)
}

func tryHttpApi(host string) {
	log.Println("Trying http api...")

	exampleApi := httpExample.New(host)

	userData, err := exampleApi.GetUserData("test")
	log.Println("http_example_res", userData, "err", err)
}

func tryCognitoUserApi(opt cognito.Options) {
	log.Println("Trying Cognito api...")

	userManager, err := cognito.NewUserManager(opt)
	if err != nil {
		log.Println("create_user_manager_err", err)
		return
	}

	userUUID, err := userManager.GetUserUUIDByEmail("test@gmail.com")
	log.Println("cognito_res", userUUID, "err", err)
}

func tryS3Storage(opt s3storage.Options) {
	log.Println("Trying s3 api...")

	fileStorage, err := s3storage.NewStorage(opt)
	if err != nil {
		log.Println("create_file_storage_err", err)
		return
	}

	s3File, err := fileStorage.Read("some_s3_url")
	log.Println("s3_storage_res", s3File, "err", err)
}
