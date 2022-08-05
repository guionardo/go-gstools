package examples

import (
	"fmt"

	"github.com/guionardo/go-gstools/cep"
)

func cep_example() {
	cep_service, err := cep.NewSQliteCEPService("test.db")
	if err != nil {
		panic(err)
	}
	cep, err := cep_service.GetCEP("89010-220")
	if err != nil {
		panic(err)
	}
	fmt.Println(cep)
}
