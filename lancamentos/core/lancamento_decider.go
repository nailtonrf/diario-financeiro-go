package lancamentos

import (
	"fluxo-go/shared/option"
	"fluxo-go/shared/result"
)

func Decide(
	l Lancamento,
	e option.Option[Lancamento],
) result.Result[LancamentoEfetuadoEvent] {

	switch v := l.(type) {

	case Credito:
		return result.Ok(
			LancamentoEfetuadoEvent{
				DadosLancamento: v.DadosLancamento,
				Tipo:            CreditoTipo,
				Motivo:          nil,
			},
		)

	case Debito:
		return result.Ok(
			LancamentoEfetuadoEvent{
				DadosLancamento: v.DadosLancamento,
				Tipo:            DebitoTipo,
				Motivo:          nil,
			},
		)

	case Estorno:
		if e.IsNone() {
			return result.Error[LancamentoEfetuadoEvent](
				ErrLancamentoOriginalNaoEncontrado,
			)
		}

		estorno := e.Unwrap()
		if estorno.Tipo() == EstornoTipo {
			return result.Error[LancamentoEfetuadoEvent](
				ErrLancamentoOriginalNaoEncontrado,
			)
		}

		dadosLancamento := v.DadosLancamento
		switch estorno.Tipo() {
		case CreditoTipo:
			dadosLancamento.Valor = dadosLancamento.Valor * -1
			dadosLancamento.Descricao = "[Estorno de crédito] - " + dadosLancamento.Descricao
		case DebitoTipo:
			dadosLancamento.Descricao = "[Estorno de débito] - " + dadosLancamento.Descricao
		}

		return result.Ok(
			LancamentoEfetuadoEvent{
				DadosLancamento: dadosLancamento,
				Tipo:            EstornoTipo,
				Motivo:          v.Motivo,
			},
		)

	default:
		return result.Error[LancamentoEfetuadoEvent](
			ErrTipoLancamentoInvalido,
		)
	}
}
