# go-api-example

[Examples](https://github.com/ksavluk/go-api-example/blob/main/cmd/main/main.go) of using various APIs:
* custom HTTP API
* AWS Cognito
* AWS S3 

##### Custom HTTP API
[api.Http()](https://github.com/ksavluk/go-api-example/blob/main/cmd/api/http.go) allows to use simplified common http operations (implemented based on "net/http") to build custom APIs (e.g. [httpExample.Api](https://github.com/ksavluk/go-api-example/tree/main/cmd/api/httpExample))

##### AWS Cognito User Manager
[cognito.UserManager](https://github.com/ksavluk/go-api-example/blob/main/cmd/api/cognito) provides operations to manage app users (register, login, logout, etc.) using AWS Cognito service.

##### AWS S3 File Storage
[s3storage.Storage](https://github.com/ksavluk/go-api-example/blob/main/cmd/api/s3storage) allows to upload and read files from AWS S3 bucket.

