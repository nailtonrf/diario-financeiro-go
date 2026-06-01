package shell

import (
	"testing"

	core "fluxo-go/lancamentos/core"
)

func TestHandleCreditoReturnsOkResponse(t *testing.T) {
	req := EfetuarLancamentoRequest{
		Body: EfetuarLancamentoBody{
			Tipo:      core.CreditoTipo,
			Valor:     200,
			Descricao: "venda",
		},
	}

	result := req.Handle()
	if result.IsError() {
		t.Fatalf("expected ok result, got error: %v", result.UnwrapError())
	}

	response := result.Unwrap()
	if response.Body.Tipo != string(core.CreditoTipo) {
		t.Fatalf("expected tipo %q, got %q", core.CreditoTipo, response.Body.Tipo)
	}
	if response.Body.Valor != 200 {
		t.Fatalf("expected valor 200, got %v", response.Body.Valor)
	}
}

func TestHandleInvalidTipoReturnsError(t *testing.T) {
	req := EfetuarLancamentoRequest{
		Body: EfetuarLancamentoBody{
			Tipo:      core.TipoLancamento("INVALIDO"),
			Valor:     10,
			Descricao: "teste",
		},
	}

	result := req.Handle()
	if result.IsOk() {
		t.Fatal("expected error result for invalid tipo")
	}
}

func TestHandleEstornoWithoutOriginalReturnsError(t *testing.T) {
	req := EfetuarLancamentoRequest{
		Body: EfetuarLancamentoBody{
			Tipo:      core.EstornoTipo,
			Valor:     10,
			Descricao: "estorno",
		},
	}

	result := req.Handle()
	if result.IsOk() {
		t.Fatal("expected error result for estorno without original")
	}
}
