package lancamentos

import "time"

type Lancamento interface {
	isLancamento()
	Tipo() TipoLancamento
	Valor() float64
}

type TipoLancamento string

const (
	CreditoTipo TipoLancamento = "CREDITO"
	DebitoTipo  TipoLancamento = "DEBITO"
	EstornoTipo TipoLancamento = "ESTORNO"
)

type DadosLancamento struct {
	ID        string
	Valor     float64
	Descricao string
	Data      time.Time
}

type Credito struct {
	DadosLancamento
}

func (Credito) isLancamento() {}

func (Credito) Tipo() TipoLancamento {
	return CreditoTipo
}

func (c Credito) Valor() float64 {
	return c.DadosLancamento.Valor
}

type Debito struct {
	DadosLancamento
}

func (Debito) isLancamento() {}

func (Debito) Tipo() TipoLancamento {
	return DebitoTipo
}

func (d Debito) Valor() float64 {
	return d.DadosLancamento.Valor
}

type Estorno struct {
	DadosLancamento
	LancamentoOriginalID string
	Motivo               *string
}

func (Estorno) isLancamento() {}

func (Estorno) Tipo() TipoLancamento {
	return EstornoTipo
}

func (e Estorno) Valor() float64 {
	return e.DadosLancamento.Valor
}

func NewCredito(dados DadosLancamento) Credito {
	return Credito{DadosLancamento: dados}
}

func NewDebito(dados DadosLancamento) Debito {
	return Debito{DadosLancamento: dados}
}

func NewEstorno(dados DadosLancamento, originalID string, motivo *string) Estorno {
	return Estorno{
		DadosLancamento:      dados,
		LancamentoOriginalID: originalID,
		Motivo:               motivo,
	}
}

type LancamentoEfetuadoEvent struct {
	DadosLancamento
	Tipo   TipoLancamento
	Motivo *string `json:"motivo,omitempty"`
}
