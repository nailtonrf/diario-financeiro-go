package shell

import core "fluxo-go/lancamentos/core"

type EfetuarLancamentoRequest struct {
	Body EfetuarLancamentoBody `json:"body"`
}

type EfetuarLancamentoBody struct {
	Tipo      core.TipoLancamento `json:"tipo"`
	Valor     float64             `json:"valor"`
	Descricao string              `json:"descricao"`

	LancamentoOriginalID *string `json:"lancamentoOriginalId,omitempty"`
	Motivo               *string `json:"motivo,omitempty"`
}
