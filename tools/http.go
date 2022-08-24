package tools

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type HttpClient interface {
	Get(url string) (resp *http.Response, err error)
}

var defaultHttpClient HttpClient

func HttpGet(requestUrl string) (body []byte, err error) {
	if defaultHttpClient == nil {
		defaultHttpClient = http.DefaultClient
	}
	response, err := defaultHttpClient.Get(requestUrl)
	if err != nil {
		return nil, err
	}
	if response.StatusCode < 100 || response.StatusCode >= 400 {
		return nil, fmt.Errorf("GET %s - status code: %d %s", requestUrl, response.StatusCode, response.Status)
	}
	defer response.Body.Close()
	return ioutil.ReadAll(response.Body)
}
