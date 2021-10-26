package httpExample

import (
	"github.com/pkg/errors"
)

const localesPath = "/locales"

type locales interface {
	GetLocales() ([]Locale, error)
}

type Locale struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

func (api *exampleApi) GetLocales() ([]Locale, error) {
	var locales []Locale
	if err := api.Get(localesPath, &locales); err != nil {
		return nil, errors.Wrap(err, "get_locales")
	}
	return locales, nil
}
