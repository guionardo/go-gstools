package cep

// https://brasilapi.com.br/

import (
	"encoding/json"
	"fmt"

	"github.com/guionardo/go-gstools/tools"
)

type (
	CEPAPIBrasilAPI struct {
		APIBase
	}
	BrasilAPIModel struct {
		CEP          string `json:"cep"`
		Street       string `json:"street"`
		Neighborhood string `json:"neighborhood"`
		State        string `json:"state"`
		City         string `json:"city"`
	}
)

func (c *CEPAPIBrasilAPI) GetCEP(cep string) (*CEP, error) {
	var model *BrasilAPIModel
	body, err := doRequest(fmt.Sprintf("https://brasilapi.com.br/api/cep/v1/%s", tools.JustNumbers(cep)))
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(body, &model); err != nil {
		return nil, err
	}

	return &CEP{
		CEP:        tools.JustNumbers(model.CEP),
		Logradouro: model.Street,
		Bairro:     model.Neighborhood,
		Municipio:  model.City,
		UF:         model.State,
	}, nil
}
