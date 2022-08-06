package cep

import (
	"fmt"
	"time"

	"github.com/guionardo/go-gstools/tools"
)

type CEPService struct {
	repository CEPRepository
}

var apis [3]CEPAPI = [3]CEPAPI{
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
	for _, api := range apis {
		cepDataReq, err := api.GetCEP(cep)
		if err == nil {
			cepDataReq.DataRequisicao = time.Now()
			err = c.repository.SaveCEP(cepDataReq)
			return cepDataReq, err
		}
	}
	if cepData == nil {
		return nil, fmt.Errorf("Nenhuma API disponível")
	}
	return cepData, nil
}
