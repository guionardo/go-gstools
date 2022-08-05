package cep

import (
	"fmt"
	"path"
	"testing"
	"time"
)

func TestNewCEPRepositorySQLite(t *testing.T) {
	t.Run("Create and populate", func(t *testing.T) {
		connectionString := fmt.Sprintf("%s", path.Join(t.TempDir(), "test.db"))
		repo, err := NewCEPRepositorySQLite(connectionString)
		if err != nil {
			t.Fatalf("NewCEPRepositorySQLite() error = %v", err)
		}
		cep := &CEP{
			CEP:            "01001000",
			Logradouro:     "Rua dos Bobos",
			TipoLogradouro: "Rua",
			Bairro:         "Bairro dos Bobos",
			Municipio:      "São Paulo",
			UF:             "SP",
			DataRequisicao: time.Now(),
		}
		if err := repo.SaveCEP(cep); err != nil {
			t.Fatalf("SaveCEP() error = %v", err)
		}
		if err := repo.SaveCEP(cep); err != nil {
			t.Fatalf("SaveCEP() error = %v", err)
		}

		cep2, err := repo.GetCEP("01001-000")
		if err != nil {
			t.Fatalf("GetCEP() error = %v", err)
		}

		if cep.Logradouro != cep2.Logradouro {
			t.Fatalf("GetCEP() = %v, want %v", cep2, cep)
		}
	})

}
