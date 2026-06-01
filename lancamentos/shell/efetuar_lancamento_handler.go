package shell

import (
	"time"

	core "fluxo-go/lancamentos/core"
	"fluxo-go/shared/option"
	"fluxo-go/shared/result"
)

func (e EfetuarLancamentoRequest) Handle() result.Result[LancamentoEfetuadoResponse] {
	return result.Map(
		result.Bind(
			e.toDomain(),
			func(l core.Lancamento) result.Result[core.LancamentoEfetuadoEvent] {
				return core.Decide(l, option.None[core.Lancamento]())
			},
		),
		toResponse,
	)
}

func (e EfetuarLancamentoRequest) toDomain() result.Result[core.Lancamento] {
	dados := core.DadosLancamento{
		Valor:     e.Body.Valor,
		Descricao: e.Body.Descricao,
		Data:      time.Now().UTC(),
	}

	switch e.Body.Tipo {
	case core.CreditoTipo:
		return result.Ok[core.Lancamento](core.NewCredito(dados))
	case core.DebitoTipo:
		return result.Ok[core.Lancamento](core.NewDebito(dados))
	case core.EstornoTipo:
		if e.Body.LancamentoOriginalID == nil {
			return result.Error[core.Lancamento](core.ErrLancamentoOriginalNaoEncontrado)
		}
		return result.Ok[core.Lancamento](core.NewEstorno(dados, *e.Body.LancamentoOriginalID, e.Body.Motivo))
	default:
		return result.Error[core.Lancamento](core.ErrTipoLancamentoInvalido)
	}
}

func toResponse(event core.LancamentoEfetuadoEvent) LancamentoEfetuadoResponse {
	return LancamentoEfetuadoResponse{
		Body: LancamentoEfetuadoBody{
			ID:        event.ID,
			Tipo:      string(event.Tipo),
			Valor:     event.Valor,
			Descricao: event.Descricao,
			Motivo:    event.Motivo,
		},
	}
}
