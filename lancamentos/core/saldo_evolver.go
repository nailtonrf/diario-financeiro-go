package lancamentos

func CalcularSaldo(inicial Saldo, lancamentos []Lancamento) Saldo {
	for _, lancamento := range lancamentos {
		switch lancamento.Tipo() {
		case CreditoTipo:
			inicial.Valor += lancamento.Valor()
		case DebitoTipo:
			inicial.Valor -= lancamento.Valor()
		case EstornoTipo:
			inicial.Valor += lancamento.Valor()
		}
	}

	return inicial
}
