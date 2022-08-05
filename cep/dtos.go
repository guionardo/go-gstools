package cep

import "time"

type CEP struct {
	CEP            string
	Logradouro     string
	TipoLogradouro string
	Bairro         string
	Municipio      string
	UF             string
	DataRequisicao time.Time
}
