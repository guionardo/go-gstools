package cep

import (
	"strings"
	"time"
)

type CEP struct {
	CEP            string    `json:"cep" yaml:"cep" toml:"cep"`
	Logradouro     string    `json:"logradouro" yaml:"logradouro" toml:"logradouro"`
	TipoLogradouro string    `json:"tipo_logradouro" yaml:"tipo_logradouro" toml:"tipo_logradouro"`
	Bairro         string    `json:"bairro" yaml:"bairro" toml:"bairro"`
	Municipio      string    `json:"municipio" yaml:"municipio" toml:"municipio"`
	UF             string    `json:"uf" yaml:"uf" toml:"uf"`
	DataRequisicao time.Time `json:"data_requisicao" yaml:"data_requisicao" toml:"data_requisicao"`
}

func (cep *CEP) CheckTipoLogradouro() {
	if len(cep.TipoLogradouro) == 0 {
		cep.TipoLogradouro = strings.ToUpper(strings.Split(cep.Logradouro, " ")[0])
	}
}
