package correios

import (
	"reflect"
	"testing"
)

func TestGetRastreio(t *testing.T) {
	type args struct {
		codigo string
	}
	tests := []struct {
		name         string
		args         args
		wantRastreio *Rastreio
		wantErr      bool
	}{
		{name: "Valid 1",
			args:         args{codigo: "LB330827204HK"},
			wantRastreio: nil,
			wantErr:      false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRastreio, err := GetRastreio(tt.args.codigo)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetRastreio() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRastreio, tt.wantRastreio) {
				t.Errorf("GetRastreio() = %v, want %v", gotRastreio, tt.wantRastreio)
			}
		})
	}
}
