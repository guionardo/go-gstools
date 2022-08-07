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

Exemplo de retorno em formato JSON

```JSON
{
 "objetos": [
  {
   "codObjeto": "LB330827204HK",
   "eventos": [
    {
     "codigo": "BDE",
     "descricao": "Objeto entregue ao destinatário",
     "dtHrCriado": "2022-05-10T13:55:02",
     "tipo": "01",
     "unidade": {
      "codSro": "",
      "endereco": {
       "cidade": "BLUMENAU",
       "uf": "SC"
      },
      "tipo": "Unidade de Distribuição"
     },
     "urlIcone": "https://proxyapp.correios.com.br/public-resources/img/smile.png"
    },
    {
     "codigo": "OEC",
     "descricao": "Objeto saiu para entrega ao destinatário",
     "dtHrCriado": "2022-05-10T07:54:12",
     "tipo": "01",
     "unidade": {
      "codSro": "",
      "endereco": {
       "bairro": "VORSTADT",
       "cep": "89015972",
       "cidade": "BLUMENAU",
       "logradouro": "RUA ANTONIO TREIS",
       "numero": "1160",
       "uf": "SC"
      },
      "tipo": "Unidade de Distribuição"
     },
     "urlIcone": "https://proxyapp.correios.com.br/public-resources/img/pre-atendimento-cor.png"
    },
    {
     "codigo": "RO",
     "descricao": "Objeto em trânsito - por favor aguarde",
     "dtHrCriado": "2022-05-09T19:55:23",
     "tipo": "01",
     "unidade": {
      "codSro": "",
      "endereco": {
       "cidade": "SAO JOSE",
       "uf": "SC"
      },
      "tipo": "Unidade de Tratamento"
     },
     "urlIcone": "https://proxyapp.correios.com.br/public-resources/img/caminhao-cor.png"
    },
    {
     "codigo": "RO",
     "descricao": "Objeto em trânsito - por favor aguarde",
     "dtHrCriado": "2022-05-06T07:57:00",
     "tipo": "01",
     "unidade": {
      "codSro": "",
      "endereco": {
       "cidade": "CURITIBA",
       "uf": "PR"
      },
      "tipo": "Unidade de Tratamento"
     },
     "urlIcone": "https://proxyapp.correios.com.br/public-resources/img/caminhao-cor.png"
    },
    {
     "codigo": "PAR",
     "descricao": "Fiscalização aduaneira finalizada",
     "dtHrCriado": "2022-05-06 07:55:00",
     "tipo": "10",
     "unidade": {
      "codSro": "",
      "endereco": {
       "cidade": "CURITIBA",
       "uf": "PR"
      },
      "tipo": "Unidade Operacional"
     },
     "urlIcone": "https://proxyapp.correios.com.br/public-resources/img/verificar-documento-cor.png"
    },
    {
     "codigo": "PAR",
     "descricao": "Objeto recebido pelos Correios do Brasil",
     "dtHrCriado": "2022-05-03T09:54:46",
     "tipo": "16",
     "unidade": {
      "codSro": "",
      "endereco": {
       "cidade": "CURITIBA",
       "uf": "PR"
      },
      "tipo": "Unidade Operacional"
     },
     "urlIcone": "https://proxyapp.correios.com.br/public-resources/img/brazil.png"
    },
    {
     "codigo": "RO",
     "descricao": "Objeto em trânsito - por favor aguarde",
     "dtHrCriado": "2022-04-26T15:09:00",
     "tipo": "01",
     "unidade": {
      "codSro": "00344000",
      "endereco": {},
      "tipo": "País",
      "nome": "HONG KONG"
     },
     "urlIcone": "https://proxyapp.correios.com.br/public-resources/img/caminhao-cor.png"
    },
    {
     "codigo": "PO",
     "descricao": "Objeto postado",
     "dtHrCriado": "2022-04-25T11:25:00",
     "tipo": "01",
     "unidade": {
      "codSro": "00344000",
      "endereco": {},
      "tipo": "País",
      "nome": "HONG KONG"
     },
     "urlIcone": "https://proxyapp.correios.com.br/public-resources/img/agencia-cor.png"
    }
   ],
   "modalidade": "V",
   "tipoPostal": {
    "categoria": "PRIME EXPRÈS",
    "descricao": "OBJETO INTERNACIONAL PRIME",
    "sigla": "LB"
   },
   "habilitaAutoDeclaracao": false,
   "permiteEncargoImportacao": false,
   "habilitaPercorridaCarteiro": false,
   "bloqueioObjeto": false,
   "possuiLocker": false,
   "habilitaLocker": false,
   "habilitaCrowdshipping": false,
   "mensagem": ""
  }
 ],
 "quantidade": 1,
 "resultado": "Todos os Eventos",
 "versao": "2.1.3"
}
```