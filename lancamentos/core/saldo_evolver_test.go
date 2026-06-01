package lancamentos

import "testing"

func TestCalcularSaldo(t *testing.T) {
	inicial := Saldo{Valor: 100}

	lancamentos := []Lancamento{
		NewCredito(DadosLancamento{Valor: 50, Descricao: "receita"}),
		NewDebito(DadosLancamento{Valor: 25, Descricao: "despesa"}),
		NewEstorno(DadosLancamento{Valor: 10, Descricao: "estorno"}, "orig-1", nil),
	}

	saldo := CalcularSaldo(inicial, lancamentos)
	if saldo.Valor != 135 {
		t.Fatalf("expected saldo 135, got %v", saldo.Valor)
	}
}
