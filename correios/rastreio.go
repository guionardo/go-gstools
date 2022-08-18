package correios

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

const url = "https://proxyapp.correios.com.br"

func GetRastreio(codigo string) (rastreio *Rastreio, err error) {
	resp, err := http.Get(fmt.Sprintf("%s/v1/sro-rastro/%s", url, codigo))
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status code: %d", resp.StatusCode)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}
	err = json.Unmarshal(body, &rastreio)

	if err == nil {
		// Ajusta urls

		for indexObj, objeto := range rastreio.Objetos {
			for indexEvt, evento := range objeto.Eventos {
				if len(evento.UrlIcone) > 0 && !strings.HasPrefix(evento.UrlIcone, "http") {
					rastreio.Objetos[indexObj].Eventos[indexEvt].UrlIcone = fmt.Sprintf("%s%s", url, evento.UrlIcone)
				}
			}
		}
	}
	return
}

func (rastreio *Rastreio) Valido() bool {
	return rastreio.Quantidade > 0 &&
		rastreio.Objetos != nil &&
		len(rastreio.Objetos[0].Eventos) > 0
}
