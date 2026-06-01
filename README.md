# Diário Financeiro Go

A financial diary (transaction ledger) application built with **Go**, implementing the **Functional Core, Imperative Shell** architecture pattern for clean separation of concerns between business logic and side effects.

## Overview

This project demonstrates best practices in Go development:
- **Functional Core**: Pure business logic in `lancamentos/core` with no side effects
- **Imperative Shell**: HTTP handlers and API layer in `lancamentos/shell`
- **Type Safety**: Leveraging Go's type system with generic `Result` and `Option` types
- **Error Handling**: Explicit error handling through `Result` types instead of panic

## Architecture

### Project Structure

```
diario-financeiro-go/
├── main.go                          # Application entry point
├── go.mod                           # Module definition
├── go.sum                           # Dependency lock file
│
├── lancamentos/
│   ├── core/                        # Functional core (pure business logic)
│   │   ├── lancamento.go            # Domain types: Lancamento, Credito, Debito, Estorno
│   │   ├── lancamento_decider.go    # Core logic for deciding transaction outcomes
│   │   ├── lancamento_errors.go     # Domain-specific errors
│   │   ├── saldo.go                 # Balance type definition
│   │   └── saldo_evolver.go         # Balance evolution logic
│   │
│   └── shell/                       # Imperative shell (HTTP handlers, I/O)
│       ├── efetuar_lancamento_handler.go    # Request handler logic
│       ├── efetuar_lancamento_request.go    # Request DTO
│       └── lancamento_efetuado_response.go  # Response DTO
│
└── shared/                          # Cross-cutting utilities
    ├── option/
    │   └── option.go                # Option type for nullable values
    └── result/
        └── result.go                # Result type for error handling
```

### Design Patterns

#### Functional Core
The `lancamentos/core` package contains pure functions with no side effects:
- **`Lancamento` interface**: Polymorphic type supporting Credito, Debito, Estorno
- **`Decide()` function**: Pure business logic that determines transaction outcomes
- **No I/O**: Core logic never touches the network, files, or databases

#### Imperative Shell
The `lancamentos/shell` package handles HTTP I/O:
- **Request handling**: Converts HTTP requests to domain objects
- **Response building**: Transforms domain results to HTTP responses
- **Side effects**: All external communication happens here

#### Generic Types
- **`Result[T]`**: Represents success/failure (`Ok` or `Error`)
- **`Option[T]`**: Represents optional values (`Some` or `None`)

## Transaction Types

### Crédito (Credit)
Represents incoming funds or positive balance adjustments.

### Débito (Debit)
Represents outgoing funds or negative balance adjustments.

### Estorno (Reversal/Refund)
Reverses a previous transaction:
- References original transaction by ID
- Automatically inverts the monetary effect
- Supports custom reversal reasons

## API Endpoints

### Create Transaction

**POST** `/lancamentos`

Create a new financial transaction.

#### Request Body
```json
{
  "tipo": "CREDITO",
  "valor": 100.50,
  "descricao": "Salary deposit",
  "lancamento_original_id": null,
  "motivo": null
}
```

**Fields:**
- `tipo` (required): `"CREDITO"`, `"DEBITO"`, or `"ESTORNO"`
- `valor` (required): Transaction amount as float64
- `descricao` (required): Transaction description
- `lancamento_original_id` (optional, required for Estorno): ID of transaction to reverse
- `motivo` (optional): Reason for reversal (Estorno only)

#### Response (Success - 200 OK)
```json
{
  "id": "uuid-string",
  "tipo": "CREDITO",
  "valor": 100.50,
  "descricao": "Salary deposit"
}
```

#### Response (Error - 400 Bad Request)
```json
{
  "title": "Invalid type",
  "status": 400,
  "detail": "Invalid transaction type"
}
```

## Getting Started

### Prerequisites
- Go 1.26.3 or later
- Git

### Installation

1. **Clone the repository**
```bash
git clone https://github.com/nailtonrf/diario-financeiro-go.git
cd diario-financeiro-go
```

2. **Download dependencies**
```bash
go mod download
```

3. **Run the server**
```bash
go run main.go
```

The API will be available at `http://localhost:8080`

### Testing the API

#### Create a credit transaction
```bash
curl -X POST http://localhost:8080/lancamentos \
  -H "Content-Type: application/json" \
  -d '{
    "tipo": "CREDITO",
    "valor": 1000.00,
    "descricao": "Monthly salary"
  }'
```

#### Create a debit transaction
```bash
curl -X POST http://localhost:8080/lancamentos \
  -H "Content-Type: application/json" \
  -d '{
    "tipo": "DEBITO",
    "valor": 250.00,
    "descricao": "Groceries"
  }'
```

#### Reverse a transaction
```bash
curl -X POST http://localhost:8080/lancamentos \
  -H "Content-Type: application/json" \
  -d '{
    "tipo": "ESTORNO",
    "valor": 250.00,
    "descricao": "Incorrect charge",
    "lancamento_original_id": "transaction-uuid-here",
    "motivo": "Duplicate charge"
  }'
```

## Dependencies

- **[Huma v2](https://github.com/danielgtaylor/huma)** (v2.38.0): REST API framework with OpenAPI support
- **[Chi v5](https://github.com/go-chi/chi)** (v5.3.0): Lightweight HTTP router

These are excellent choices for Go web development:
- Huma provides automatic OpenAPI documentation
- Chi offers composable middleware and routing

## Code Quality & Best Practices

### Error Handling
```go
// Result type prevents panic and forces error handling
result := core.Decide(lancamento, option)
if result.IsError() {
    return nil, result.UnwrapError()
}
value := result.Unwrap()
```

### Type Safety
- No `interface{}` for core logic
- Generics prevent nil pointer bugs
- Sealed interface pattern (`isLancamento()`) prevents external implementations

### Separation of Concerns
- **Core**: Pure, testable, no dependencies
- **Shell**: Thin adapter layer for HTTP concerns
- **Shared**: Reusable utilities

## Future Improvements

- [ ] Persistence layer (database integration)
- [ ] Balance calculation and reporting
- [ ] Transaction querying and filtering
- [ ] User authentication and authorization
- [ ] Unit tests and integration tests
- [ ] Docker containerization
- [ ] Comprehensive API logging
- [ ] Transaction validation rules engine
- [ ] Event sourcing for audit trail

## Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feat/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feat/amazing-feature`)
5. Open a Pull Request

## License

This project is open source. See LICENSE file for details.

## Author

Created by [nailtonrf](https://github.com/nailtonrf)

---

**Built with functional programming principles and Go best practices** 🚀
