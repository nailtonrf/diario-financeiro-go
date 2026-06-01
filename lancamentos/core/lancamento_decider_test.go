package lancamentos

import (
	"testing"
	"time"

	"fluxo-go/shared/option"
)

func TestDecideCreditoOk(t *testing.T) {
	dados := DadosLancamento{Valor: 100, Descricao: "receita", Data: time.Now()}
	credito := NewCredito(dados)

	result := Decide(credito, option.None[Lancamento]())
	if result.IsError() {
		t.Fatalf("expected ok result, got error: %v", result.UnwrapError())
	}

	event := result.Unwrap()
	if event.Tipo != CreditoTipo {
		t.Fatalf("expected tipo %v, got %v", CreditoTipo, event.Tipo)
	}
	if event.Valor != 100 {
		t.Fatalf("expected valor 100, got %v", event.Valor)
	}
}

func TestDecideEstornoCreditoOriginal(t *testing.T) {
	dados := DadosLancamento{Valor: 50, Descricao: "estorno", Data: time.Now()}
	estorno := NewEstorno(dados, "orig-1", nil)

	originalDados := DadosLancamento{Valor: 100, Descricao: "compra", Data: time.Now()}
	var original Lancamento = NewCredito(originalDados)

	result := Decide(estorno, option.Some(original))
	if result.IsError() {
		t.Fatalf("expected ok result, got error: %v", result.UnwrapError())
	}

	event := result.Unwrap()
	if event.Tipo != EstornoTipo {
		t.Fatalf("expected tipo %v, got %v", EstornoTipo, event.Tipo)
	}
	if event.Valor != -50 {
		t.Fatalf("expected valor -50, got %v", event.Valor)
	}
	if event.Descricao != "[Estorno de crédito] - estorno" {
		t.Fatalf("expected descricao modified, got %q", event.Descricao)
	}
}

func TestDecideEstornoSemOriginalRetornaErro(t *testing.T) {
	dados := DadosLancamento{Valor: 10, Descricao: "estorno", Data: time.Now()}
	estorno := NewEstorno(dados, "orig-1", nil)

	result := Decide(estorno, option.None[Lancamento]())
	if result.IsOk() {
		t.Fatal("expected error result for missing original lancamento")
	}
}
