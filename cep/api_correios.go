package cep

import (
	"encoding/json"
	"fmt"

	"github.com/guionardo/go-gstools/tools"
)

/*
	{
	    "erro": false,
	    "mensagem": "DADOS ENCONTRADOS COM SUCESSO.",
	    "total": 1,
	    "dados": [
	        {
	            "uf": "SC",
	            "localidade": "Blumenau",
	            "locNoSem": "",
	            "locNu": "",
	            "localidadeSubordinada": "",
	            "logradouroDNEC": "Rua Manoel Barreto",
	            "logradouroTextoAdicional": "",
	            "logradouroTexto": "",
	            "bairro": "Victor Konder",
	            "baiNu": "",
	            "nomeUnidade": "",
	            "cep": "89012134",
	            "tipoCep": "2",
	            "numeroLocalidade": "",
	            "situacao": "",
	            "faixasCaixaPostal": [],
	            "faixasCep": []
	        }
	    ]
	}
*/
type (
	CEPCorreiosAPI struct {
	}
	CorreiosAPIModel struct {
		UF                       string `json:"uf"`
		Localidade               string `json:"localidade"`
		LogradouroDNEC           string `json:"logradouroDNEC"`
		LogradouroTextoAdicional string `json:"logradouroTextoAdicional"`
		LogradouroTexto          string `json:"logradouroTexto"`
		Bairro                   string `json:"bairro"`
		CEP                      string `json:"cep"`

		Street       string `json:"street"`
		Neighborhood string `json:"neighborhood"`
		State        string `json:"state"`
		City         string `json:"city"`
	}
	CorreiosAPIModelResponse struct {
		Erro     bool               `json:"erro"`
		Mensagem string             `json:"mensagem"`
		Total    int                `json:"total"`
		Dados    []CorreiosAPIModel `json:"dados"`
	}
)

func (c *CEPCorreiosAPI) GetCEP(cep string) (*CEP, error) {
	data := map[string]string{
		"endereco": tools.JustNumbers(cep),
		"tipoCEP":  "ALL",
	}
	body, err := tools.HttpPostUrlEncoded("https://buscacepinter.correios.com.br/app/endereco/carrega-cep-endereco.php", data)
	if err != nil {
		return nil, err
	}
	var model *CorreiosAPIModelResponse
	if err := json.Unmarshal(body, &model); err != nil {
		return nil, err
	}
	if model.Erro || len(model.Dados) == 0 {
		return nil, fmt.Errorf("CEP %s não encontrado (%s)", cep, model.Mensagem)
	}
	cepResponse := model.Dados[0]
	return &CEP{
		CEP:        tools.JustNumbers(cepResponse.CEP),
		Logradouro: cepResponse.LogradouroDNEC,
		Bairro:     cepResponse.Bairro,
		Municipio:  cepResponse.Localidade,
		UF:         cepResponse.UF,
	}, nil

}
