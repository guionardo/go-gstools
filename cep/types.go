package cep

type (
	CEPAPIProvider interface {
		GetCEP(cep string) (*CEP, error)
	}

	CEPRepository interface {
		GetCEP(cep string) (*CEP, error)
		SaveCEP(cep *CEP) error
	}
)
