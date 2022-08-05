package examples

import (
	"encoding/json"
	"fmt"

	"github.com/guionardo/go-gstools/correios"
)

func main() {
	rastreio, err := correios.GetRastreio("LB330827204HK")
	if err != nil {

		fmt.Printf("Erro na consulta aos correios: %v", err)
	} else {
		jsonRastreio, err := json.MarshalIndent(rastreio, "", " ")
		if err != nil {
			fmt.Printf("Erro ao gerar json: %v", err)
		} else {
			fmt.Printf("%s", jsonRastreio)
		}
	}
}
