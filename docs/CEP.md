## CEP

Pacote de consulta a API's grátis de CEP e persistência em cache

```GoLang
package main

import (
	"fmt"

	"github.com/guionardo/go-gstools/cep"
)

func main() {
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
```