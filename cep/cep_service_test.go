package cep

import (
	"path"
	"testing"
)

func TestNewCEPService(t *testing.T) {
	t.Run("CEPService + SQLite", func(t *testing.T) {
		repo, err := NewCEPRepositorySQLite(path.Join(t.TempDir(), "test.db"))
		if err != nil {
			t.Fatalf("NewCEPRepositorySQLite() error = %v", err)
		}
		service := NewCEPService(repo)
		cep, err := service.GetCEP("89010-220")
		if err != nil {
			t.Fatalf("GetCEP() error = %v", err)
		}
		if cep.Logradouro != "Rua São José" {
			t.Fatalf("GetCEP() = %v, want %v", cep, "Rua São José")
		}
		cep2, err := service.GetCEP("89010-220")
		if err != nil {
			t.Fatalf("GetCEP() error = %v", err)
		}
		if cep.Logradouro != cep2.Logradouro || cep.CEP != cep2.CEP {
			t.Fatalf("GetCEP() = %v, want %v", cep2, cep)
		}

	})

}
