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
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}
	err = json.Unmarshal(body, &rastreio)

	if err == nil {
		// Ajusta urls

		for _, objeto := range rastreio.Objetos {
			for _, evento := range objeto.Eventos {
				if len(evento.UrlIcone) > 0 && !strings.HasPrefix(evento.UrlIcone, "http") {
					evento.UrlIcone = url + evento.UrlIcone
				}
			}
		}
	}
	return
}
