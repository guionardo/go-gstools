package cep

import (
	"reflect"
	"testing"
)

func TestCEPAPIViaCep_GetCEP(t *testing.T) {
	tests := []struct {
		name string
		cep  string
		wantLogradouro string
		wantErr        bool
	}{
		{
			name:           "89010220",
			cep:            "89010220",
			wantLogradouro: "Rua São José",
			wantErr:        false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &CEPAPIViaCep{}

			got, err := c.GetCEP(tt.cep)
			if (err != nil) != tt.wantErr {
				t.Errorf("CEPAPIViaCep.GetCEP() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got.Logradouro, tt.wantLogradouro) {
				t.Errorf("CEPAPIViaCep.GetCEP() = %v, want %v", got.Logradouro, tt.wantLogradouro)
			}
		})
	}
}
