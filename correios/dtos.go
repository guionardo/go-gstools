package correios

import "github.com/guionardo/go-gstools/tools"

type (
	Rastreio struct {
		Objetos    []Objeto `json:"objetos"`
		Quantidade int      `json:"quantidade"`
		Resultado  string   `json:"resultado"`
		Versao     string   `json:"versao"`
	}
	Objeto struct {
		CodObjeto                  string     `json:"codObjeto"`
		Eventos                    []Evento   `json:"eventos"`
		Modalidade                 string     `json:"modalidade"`
		TipoPostal                 TipoPostal `json:"tipoPostal"`
		HabilitaAutoDeclaracao     bool       `json:"habilitaAutoDeclaracao"`
		PermiteEncargoImportacao   bool       `json:"permiteEncargoImportacao"`
		HabilitaPercorridaCarteiro bool       `json:"habilitaPercorridaCarteiro"`
		BloqueioObjeto             bool       `json:"bloqueioObjeto"`
		PossuiLocker               bool       `json:"possuiLocker"`
		HabilitaLocker             bool       `json:"habilitaLocker"`
		HabilitaCrowdShipping      bool       `json:"habilitaCrowdshipping"`
		Mensagem                   string     `json:"mensagem,omitempty"`
	}
	Evento struct {
		Codigo         string     `json:"codigo"`
		Descricao      string     `json:"descricao"`
		DataCriado     CustomTime `json:"dtHrCriado"`
		Tipo           string     `json:"tipo"`
		Unidade        Unidade    `json:"unidade"`
		UnidadeDestino Unidade    `json:"unidadeDestino"`
		UrlIcone       string     `json:"urlIcone"`
	}
	TipoPostal struct {
		Categoria string `json:"categoria"`
		Descricao string `json:"descricao"`
		Sigla     string `json:"sigla"`
	}
	Unidade struct {
		CodSro   string   `json:"codSro"`
		Endereco Endereco `json:"endereco,omitempty"`
		Tipo     string   `json:"tipo"`
		Nome     string   `json:"nome,omitempty"`
	}
	Endereco struct {
		Bairro     string `json:"bairro,omitempty"`
		Cep        string `json:"cep,omitempty"`
		Cidade     string `json:"cidade,omitempty"`
		Logradouro string `json:"logradouro,omitempty"`
		Numero     string `json:"numero,omitempty"`
		UF         string `json:"uf,omitempty"`
	}
)

func (um *Unidade) String() string {
	return tools.JoinNotEmpty(um.Tipo, um.Nome, um.CodSro, um.Endereco.String())
}
func (em *Endereco) String() string {
	return tools.JoinNotEmpty(em.Logradouro, em.Numero, em.Bairro, em.Cidade, em.UF)
}
