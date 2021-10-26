package api

import (
	"bytes"
	"encoding/json"
	"github.com/pkg/errors"
	"io"
	"io/ioutil"
	"net/http"
)

type HttpApiOptions struct {
	Host string
}

type HttpApi interface {
	Get(path string, result interface{}) error
	Post(path string, request interface{}, result interface{}) error
	Delete(path string, result interface{}) error
}

type httpApi struct {
	host string
}

func Http(opt HttpApiOptions) HttpApi {
	return &httpApi{
		host: opt.Host,
	}
}

func (api *httpApi) Get(path string, result interface{}) error {
	resp, err := http.Get(api.host + path)
	if err != nil {
		return errors.Wrap(err, "api_do_get_request")
	}
	defer resp.Body.Close()

	return api.readResponse(resp, &result)
}

func (api *httpApi) Post(path string, request interface{}, result interface{}) error {
	requestReader, err := api.getReader(request)
	if err != nil {
		return errors.Wrap(err, "api_read_request")
	}

	httpRequest, err := http.NewRequest(http.MethodPost, api.host+path, requestReader)
	if err != nil {
		return errors.Wrap(err, "api_create_post_request")
	}

	resp, err := http.DefaultClient.Do(httpRequest)
	if err != nil {
		return errors.Wrap(err, "api_do_post_request")
	}
	defer resp.Body.Close()

	return api.readResponse(resp, &result)
}

func (api *httpApi) Delete(path string, result interface{}) error {
	httpRequest, err := http.NewRequest(http.MethodDelete, api.host+path, nil)
	if err != nil {
		return errors.Wrap(err, "api_create_delete_request")
	}

	resp, err := http.DefaultClient.Do(httpRequest)
	if err != nil {
		return errors.Wrap(err, "api_do_delete_request")
	}
	defer resp.Body.Close()

	return api.readResponse(resp, &result)
}

func (api *httpApi) readResponse(resp *http.Response, result interface{}) error {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return errors.Wrap(err, "api_read_response")
	}

	if err := json.Unmarshal(body, &result); err != nil {
		return errors.Wrap(err, "api_unmarshal_response")
	}

	return nil
}

func (api *httpApi) getReader(body interface{}) (io.Reader, error) {
	if body == nil {
		return nil, nil
	}
	json, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	return bytes.NewReader(json), nil
}
