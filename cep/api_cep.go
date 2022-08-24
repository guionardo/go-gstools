package cep

import (
	"encoding/json"
	"fmt"

	"github.com/guionardo/go-gstools/tools"
)

type (
	CEPAPICep struct {
	}
	ApiCepModel struct {
		Code       string `json:"code"`
		Address    string `json:"address"`
		District   string `json:"district"`
		City       string `json:"city"`
		State      string `json:"state"`
		StatusText string `json:"statusText"`
		Status     int    `json:"status"`
		Ok         bool   `json:"ok"`
		Message    string `json:"message"`
	}
)

func (c *CEPAPICep) GetCEP(cep string) (*CEP, error) {
	var model *ApiCepModel
	body, err := tools.HttpGet(fmt.Sprintf("https://ws.apicep.com/cep/%s.json", tools.JustNumbers(cep)))
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(body, &model); err != nil {
		return nil, err
	}
	if model.Ok && model.Status == 200 {
		return &CEP{
			CEP:        tools.JustNumbers(cep),
			Logradouro: model.Address,
			Bairro:     model.District,
			Municipio:  model.City,
			UF:         model.State,
		}, nil
	}
	return nil, fmt.Errorf("CEP %s não encontrado (%d %v)", cep, model.Status, model.Message)
}
