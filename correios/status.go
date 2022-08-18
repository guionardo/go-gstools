package correios

type Status struct {
	Icon   string
	Text   string
	Detail string
}

var Statuses = map[string]Status{
	"BDE": {
		Icon: "🏠",
		Text: "Objeto entregue ao destinatário",
	},
	"OEC": {
		Icon: "📨",
		Text: "Objeto saiu para entrega ao destinatário",
	},
	"FC": {
		Icon:   "🛣️",
		Text:   "Objeto em correção de rota",
		Detail: "Corrigimos um equívoco no encaminhamento do seu objeto. Por favor aguarde",
	},
	"PO": {
		Icon: "🏤",
		Text: "Objeto postado",
	},
	"RO": {
		Icon: "🚚",
		Text: "Objeto em trânsito - por favor aguarde",
	},
	"PAR": {
		Icon: "🏢",
		Text: "Objeto recebido pelos Correios do Brasil",
	},
	"DO":{
		Icon: "🚚",
		Text: "Objeto em trânsito - por favor aguarde",
	},
}
