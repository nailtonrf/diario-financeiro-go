package lancamentos

import (
	core "fluxo-go/lancamentos/core"
	"fluxo-go/shared/option"
	"fluxo-go/shared/result"
	"time"
)

func (e EfetuarLancamentoRequest) Handle() result.Result[LancamentoEfetuadoResponse] {

	dados := core.DadosLancamento{
		Valor:     e.Body.Valor,
		Descricao: e.Body.Descricao,
		Data:      time.Now().UTC(),
	}

	var lancamento core.Lancamento

	switch e.Body.Tipo {

	case core.CreditoTipo:

		lancamento = core.Credito{
			DadosLancamento: dados,
		}

	case core.DebitoTipo:

		lancamento = core.Debito{
			DadosLancamento: dados,
		}

	case core.EstornoTipo:

		if e.Body.LancamentoOriginalID == nil {

			return result.Error[LancamentoEfetuadoResponse](
				core.
					ErrLancamentoOriginalNaoEncontrado,
			)
		}

		lancamento = core.Estorno{

			DadosLancamento: dados,

			LancamentoOriginalID: *e.Body.LancamentoOriginalID,

			Motivo: deref(
				e.Body.Motivo,
			),
		}

	default:

		return result.Error[LancamentoEfetuadoResponse](
			core.
				ErrTipoLancamentoInvalido,
		)
	}

	deciderResult :=
		core.
			Decide(
				lancamento,
				option.None[core.Lancamento](),
			)

	if deciderResult.IsError() {

		return result.Error[LancamentoEfetuadoResponse](
			deciderResult.UnwrapError(),
		)
	}

	event :=
		deciderResult.Unwrap()

	response :=
		LancamentoEfetuadoResponse{

			Body: LancamentoEfetuadoBody{

				ID: event.ID,

				Tipo: string(
					event.Tipo,
				),

				Valor: event.Valor,

				Descricao: event.Descricao,
			},
		}

	return result.Ok(
		response,
	)
}

func deref(
	value *string,
) string {

	if value == nil {
		return ""
	}

	return *value
}
