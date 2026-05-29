package main

import (
	"context"
	"net/http"
	"time"

	lancamentos "fluxo-go/lancamentos/core"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humachi"
	"github.com/go-chi/chi/v5"
)

func main() {

	router := chi.NewRouter()

	api := humachi.New(
		router,
		huma.DefaultConfig(
			"Fluxo API",
			"1.0.0",
		),
	)

	huma.Get(
		api,
		"/lancamentos",
		func(
			ctx context.Context,
			input *struct{},
		) (*ListarLancamentosResponse, error) {

			response := ListarLancamentosResponse{
				Body: []lancamentos.Credito{
					{
						DadosLancamento: lancamentos.DadosLancamento{
							ID:        "1",
							Valor:     150.00,
							Descricao: "Depósito",
							Data:      time.Now(),
						},
					},
				},
			}

			return &response, nil
		},
	)

	http.ListenAndServe(
		":8080",
		router,
	)
}

type ListarLancamentosResponse struct {
	Body []lancamentos.Credito `json:"body"`
}
