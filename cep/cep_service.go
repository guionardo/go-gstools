package cep

import (
	"fmt"
	"time"

	"github.com/guionardo/go-gstools/tools"
)

type CEPService struct {
	repository CEPRepository
}

var apis [2]CEPAPI = [2]CEPAPI{
	&CEPAPIViaCep{},
	&CEPAPICep{},
}

func NewCEPService(repository CEPRepository) *CEPService {
	return &CEPService{
		repository: repository,
	}
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
