package lancamentos

func (s Saldo) SaldoAtual(p Saldo, l []Lancamento) Saldo {

	for _, lancamento := range l {
		if lancamento.Tipo() == CreditoTipo {
			p.Valor += lancamento.Valor()
		}
		if lancamento.Tipo() == DebitoTipo {
			p.Valor -= lancamento.Valor()
		}
		if lancamento.Tipo() == EstornoTipo {
			p.Valor += lancamento.Valor()
		}
	}

	return p
}
