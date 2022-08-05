# go-gstools

GoLang Guiosoft Tools

[![Go](https://github.com/guionardo/go-gstools/actions/workflows/go.yml/badge.svg)](https://github.com/guionardo/go-gstools/actions/workflows/go.yml)

## Rastreamento correios

```GoLang
package main

import (
	"fmt"

	"github.com/guionardo/go-gstools/correios"
)

func main() {
	rastreio, err := correios.GetRastreio("LB330827204HK")
	if err != nil {
		fmt.Printf("Erro na consulta aos correios: %v", err)
	} else {
		fmt.Printf("%v", rastreio)
	}
}
```