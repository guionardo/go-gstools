package cep

import (
	"fmt"
	"time"

	"github.com/guionardo/go-gstools/tools"
)

type CEPService struct {
	repository CEPRepository
}

var apis [3]CEPAPIProvider = [3]CEPAPIProvider{
	&CEPAPIViaCep{},
	&CEPAPICep{},
	&CEPAPIBrasilAPI{},
}

func NewCEPService(repository CEPRepository) *CEPService {
	return &CEPService{
		repository: repository,
	}
}

func NewSQliteCEPService(connectionString string) (*CEPService, error) {
	repo, err := NewCEPRepositorySQLite(connectionString)
	if err == nil {
		return NewCEPService(repo), nil
	}
	return nil, err
}

func (c *CEPService) GetCEP(cep string) (*CEP, error) {
	cep = tools.JustNumbers(cep)
	cepData, err := c.repository.GetCEP(cep)
	if err == nil && cepData != nil && cepData.DataRequisicao.After(time.Now().Add(-24*time.Hour*30)) {
		return cepData, nil
	}
	var errors []error
	ch := make(chan *CEP)
	errorChan := make(chan error)
	for _, api := range apis {
		go func(api CEPAPIProvider) {
			cepDataReq, err := api.GetCEP(cep)
			if err == nil {
				ch <- cepDataReq
			} else {
				errorChan <- err
			}
		}(api)
	}
	for i := 0; i < len(apis); i++ {
		select {
		case res := <-ch:
			if res != nil {
				res.DataRequisicao = time.Now()
				res.CheckTipoLogradouro()
				err = c.repository.SaveCEP(res)
				return res, err
			}
		case err := <-errorChan:
			if err != nil {
				errors = append(errors, err)
			}
		}
	}
	if len(errors) == len(apis) {
		return nil, fmt.Errorf("Nenhuma API disponível")
	}
	return nil, fmt.Errorf("Nenhuma API retornou dados")
}
