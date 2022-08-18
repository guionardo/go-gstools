package cep

// https://viacep.com.br/

import (
	"encoding/json"
	"fmt"

	"github.com/guionardo/go-gstools/tools"
)

type (
	CEPAPIViaCep struct {
		APIBase
	}
	ViaCepModel struct {
		CEP         string `json:"cep"`
		Logradouro  string `json:"logradouro"`
		Complemento string `json:"complemento"`
		Bairro      string `json:"bairro"`
		Localidade  string `json:"localidade"`
		UF          string `json:"uf"`
		IBGE        string `json:"ibge"`
		GIA         string `json:"gia"`
		DDD         string `json:"ddd"`
		SIAFI       string `json:"siafi"`
		ERRO        string `json:"erro"`
	}
)

func (c *CEPAPIViaCep) GetCEP(cep string) (*CEP, error) {
	var model *ViaCepModel
	body, err := doRequest(fmt.Sprintf("https://viacep.com.br/ws/%s/json/", tools.JustNumbers(cep)))
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(body, &model); err != nil {
		return nil, err
	}
	if len(model.ERRO) > 0 {
		return nil, fmt.Errorf("CEP %s não encontrado", cep)
	}
	return &CEP{
		CEP:        tools.JustNumbers(model.CEP),
		Logradouro: model.Logradouro,
		Bairro:     model.Bairro,
		Municipio:  model.Localidade,
		UF:         model.UF,
	}, nil
}
