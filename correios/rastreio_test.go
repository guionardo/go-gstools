package correios

import (
	"testing"
)

func TestGetRastreio(t *testing.T) {
	type args struct {
		codigo string
	}
	tests := []struct {
		name      string
		codigo    string
		args      args
		mensagem  string
		wantErr   bool
		wantValid bool
	}{
		{
			name:      "Valid 1",
			codigo:    "NA680830120BR",
			mensagem:  "",
			wantErr:   false,
			wantValid: true,
		}, {
			name:      "Invalid 1",
			codigo:    "NANANANANA",
			mensagem:  "SRO-019: Objeto inválido",
			wantErr:   false,
			wantValid: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRastreio, err := GetRastreio(tt.codigo)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetRastreio() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotRastreio.Quantidade != 1 {
				t.Errorf("GetRastreio() = %v, want 1", gotRastreio)
			}
			if gotRastreio.Objetos[0].Mensagem != tt.mensagem {
				t.Errorf("GetRastreio() = %v, want %v", gotRastreio.Objetos[0].Mensagem, tt.mensagem)
			}
			if tt.wantValid != gotRastreio.Valido() {
				t.Errorf("GetRastreio().Valido() = %v, want %v", gotRastreio.Valido(), tt.wantValid)
			}
		})
	}
}
