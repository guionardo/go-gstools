package cep

import (
	"strings"
	"testing"
)

func TestCEPAPIs_GetCEP(t *testing.T) {
	cep := "89010220"
	wantLogradouro := "Rua São José"
	tests := []struct {
		name string
		api  CEPAPI
	}{
		{
			name: "BrasilAPI",
			api:  &CEPAPIBrasilAPI{},
		},
		{
			name: "APICEP",
			api:  &CEPAPICep{},
		},
		{
			name: "ViaCEP",
			api:  &CEPAPIViaCep{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.api.GetCEP("00000000")
			if err == nil {
				t.Errorf("CEPAPI.GetCEP() [%s] must return error for unexistent CEP", tt.name)
				return
			}
			got, err = tt.api.GetCEP(cep)
			if err != nil {
				if strings.Contains(err.Error(), "flood") ||
					strings.Contains(err.Error(), "bad_request") {
					t.Logf("CEPAPI.GetCEP() [%s] error = %v", tt.name, err)
					return
				}
				t.Errorf("CEPAPI.GetCEP() [%s] error = %v", tt.name, err)
				return
			}
			if got.CEP != cep {
				t.Errorf("CEPAPI.GetCEP() [%s] = %v, want %v", tt.name, got, cep)
			}
			if got.Logradouro != wantLogradouro {
				t.Errorf("CEPAPI.GetCEP() [%s] = %v, want %v", tt.name, got.Logradouro, wantLogradouro)
			}

		})
	}
}
