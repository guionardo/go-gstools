package cep

import (
	"testing"
	"time"
)

func TestCEP_CheckTipoLogradouro(t *testing.T) {
	type fields struct {
		CEP            string
		Logradouro     string
		TipoLogradouro string
		Bairro         string
		Municipio      string
		UF             string
		DataRequisicao time.Time
	}
	tests := []struct {
		name     string
		cep      *CEP
		wantTipo string
	}{
		{
			name:     "Valid 1",
			cep:      &CEP{Logradouro: "Rua das Flores"},
			wantTipo: "RUA",
		},
		{
			name:     "Empty 1",
			cep:      &CEP{},
			wantTipo: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.cep.CheckTipoLogradouro()
			if tt.cep.TipoLogradouro != tt.wantTipo {
				t.Errorf("CEP.CheckTipoLogradouro() = %v, want %v", tt.cep.TipoLogradouro, tt.wantTipo)
			}
		})
	}
}
