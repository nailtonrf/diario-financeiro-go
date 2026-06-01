package main

import (
	"context"
	"net/http"

	shell "fluxo-go/lancamentos/shell"

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

	huma.Post(
		api,
		"/lancamentos",

		func(
			ctx context.Context,
			input *shell.
				EfetuarLancamentoRequest,
		) (
			*shell.LancamentoEfetuadoResponse,
			error,
		) {

			result :=
				input.Handle()

			if result.IsError() {
				return nil, result.UnwrapError()
			}

			response := result.Unwrap()

			return &response, nil
		},
	)

	http.ListenAndServe(
		":8080",
		router,
	)
}
