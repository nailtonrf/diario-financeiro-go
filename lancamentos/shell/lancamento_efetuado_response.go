package shell

type LancamentoEfetuadoResponse struct {
	Body LancamentoEfetuadoBody `json:"body"`
}

type LancamentoEfetuadoBody struct {
	ID        string  `json:"id"`
	Tipo      string  `json:"tipo"`
	Valor     float64 `json:"valor"`
	Descricao string  `json:"descricao"`
	Motivo    *string `json:"motivo,omitempty"`
}
