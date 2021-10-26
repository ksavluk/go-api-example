package httpExample

import (
	"ksavluk/go-api-example/cmd/api"
)

type Api interface {
	locales
	users
}

type exampleApi struct {
	api.HttpApi
}

func New(host string) Api {
	opt := api.HttpApiOptions{Host: host}
	return &exampleApi{
		HttpApi: api.Http(opt),
	}
}
