# Diário Financeiro Go - Contributing Guide

## Development Setup

### Prerequisites
- Go 1.26.3 or later
- Git

### Local Development

1. **Clone and setup**
```bash
git clone https://github.com/nailtonrf/diario-financeiro-go.git
cd diario-financeiro-go
go mod download
```

2. **Run locally**
```bash
go run main.go
```

3. **Test API**
```bash
curl -X POST http://localhost:8080/lancamentos \
  -H "Content-Type: application/json" \
  -d '{"tipo":"CREDITO","valor":100,"descricao":"Test"}'
```

## Project Guidelines

### Architecture Principles

- **Functional Core**: `lancamentos/core/` contains pure functions with deterministic behavior
- **Imperative Shell**: `lancamentos/shell/` handles HTTP I/O and side effects
- **Type Safety**: Leverage Go's type system; avoid `interface{}` in core logic

### Code Organization

- **Domain Logic**: Place in `lancamentos/core/`
- **HTTP Handlers**: Place in `lancamentos/shell/`
- **Shared Utilities**: Place in `shared/`
- **DTOs**: Keep request/response types in `shell/`

### Naming Conventions

- **Portuguese domain terms**: Use Portuguese for business concepts (lancamento, credito, debito, estorno)
- **PascalCase for types**: `Lancamento`, `Credito`, `Debito`, `Estorno`
- **camelCase for functions**: `Handle()`, `Decide()`, `Unwrap()`
- **File naming**: Use snake_case with domain suffixes (e.g., `lancamento_decider.go`, `efetuar_lancamento_handler.go`)

### Error Handling

Always use `Result[T]` for functions that may fail:

```go
func ProcessTransaction(lancamento Lancamento) Result[LancamentoEfetuadoEvent] {
    // Never panic in core logic
    // Return Error[T] instead
}
```

### Testing

When adding tests:
- Create `*_test.go` files in the same package
- Test core logic thoroughly
- Keep shell tests focused on HTTP concerns

## Commit Message Format

Follow conventional commits:

```
type: subject

body

footer
```

**Types:**
- `feat`: New feature
- `fix`: Bug fix
- `docs`: Documentation
- `refactor`: Code refactoring
- `test`: Adding tests
- `chore`: Build, dependencies, etc.

**Examples:**
```
feat: add transaction filtering by date range

docs: update API documentation with examples

fix: handle nil pointers in deref function
```

## Pull Request Process

1. Create feature branch from `main`: `git checkout -b feat/your-feature`
2. Make commits following conventional format
3. Ensure code follows project style
4. Write meaningful PR description
5. Request review from project maintainers

## Future Work

- [ ] Add unit tests with table-driven tests
- [ ] Add integration tests
- [ ] Add database persistence layer
- [ ] Add balance calculation endpoints
- [ ] Add transaction history queries
- [ ] Add authentication/authorization
- [ ] Set up CI/CD pipeline
- [ ] Docker and docker-compose setup

## Questions or Issues?

Open an issue on GitHub describing:
- What you're trying to do
- Expected behavior
- Actual behavior
- Steps to reproduce (for bugs)
- Your environment (Go version, OS, etc.)
