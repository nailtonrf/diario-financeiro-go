package lancamentos

import "errors"

var ErrTipoLancamentoInvalido = errors.New("tipo de lançamento inválido")
var ErrLancamentoOriginalNaoEncontrado = errors.New("lançamento original para estorno não encontrado")
