package cep

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type (
	CEPAPI interface {
		GetCEP(cep string) (*CEP, error)
	}
	APIBase struct {
	}
	CEPRepository interface {
		GetCEP(cep string) (*CEP, error)
		SaveCEP(cep *CEP) error
	}
)

func doRequest(requestUrl string) (body []byte, err error) {
	response, err := http.DefaultClient.Get(requestUrl)
	if err != nil {
		return nil, err
	}
	if response.StatusCode < 100 || response.StatusCode >= 400 {
		return nil, fmt.Errorf("GET %s - status code: %d %s", requestUrl, response.StatusCode, response.Status)
	}
	defer response.Body.Close()
	return ioutil.ReadAll(response.Body)
}
