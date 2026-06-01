# Architecture Decision Records (ADRs)

## ADR-001: Functional Core, Imperative Shell Pattern

**Status**: Accepted

**Context**: The project needs a clear separation between business logic and I/O concerns.

**Decision**: Implement Functional Core, Imperative Shell pattern:
- **Core** (`lancamentos/core/`): Pure functions with no side effects, deterministic behavior
- **Shell** (`lancamentos/shell/`): HTTP handlers, request/response transformation, I/O

**Rationale**:
- Easy to test core logic without mocking I/O
- Clear dependency direction (Shell depends on Core, not vice versa)
- Simple to reason about control flow
- Facilitates future persistence layer changes

**Consequences**:
- Core package cannot directly call HTTP endpoints or database
- Shell layer must handle all external communication
- Clear interface contracts between layers

---

## ADR-002: Generic Result and Option Types

**Status**: Accepted

**Context**: Go lacks built-in Result and Option types. Error handling typically uses `(T, error)` tuples or panics.

**Decision**: Implement generic `Result[T]` and `Option[T]` types in `shared/result` and `shared/option`.

**Rationale**:
- Prevent nil pointer panics
- Force explicit error handling at call sites
- Enable functional programming patterns (Map, Bind, Ensure)
- More expressive than bare error returns

**Consequences**:
- Slightly more verbose at call sites
- But prevents entire classes of runtime errors
- Monadic composition enables elegant error chains

**Example**:
```go
// Before: Easy to forget error handling
result, _ := risky()

// After: Error handling is explicit
result := risky()
if result.IsError() {
    return result.UnwrapError()
}
```

---

## ADR-003: Transaction Types as Sealed Interface

**Status**: Accepted

**Context**: Need type-safe way to represent different transaction kinds without allowing external implementations.

**Decision**: Use sealed interface pattern with private `isLancamento()` method.

**Rationale**:
- Compiler enforces exhaustive switch statements
- Prevents external code from implementing interface
- Clear domain model
- Type-safe pattern matching

**Consequences**:
- New transaction types require code changes in core
- Can't extend behavior from external packages
- Forces intentional API changes for new capabilities

---

## ADR-004: Portuguese Domain Language

**Status**: Accepted

**Context**: Project is a financial diary (diário financeiro) for a Portuguese-speaking audience.

**Decision**: Use Portuguese terminology for domain concepts:
- `Lancamento` (transaction)
- `Credito` (credit)
- `Debito` (debit)
- `Estorno` (reversal/refund)
- `Saldo` (balance)

**Rationale**:
- Domain experts think in Portuguese
- Reduces translation errors in requirements
- Improves communication with stakeholders
- Clear domain-driven design

**Consequences**:
- Go naming conventions differ from typical English codebases
- Easier for Portuguese-speaking team members
- External contributors may need glossary

---

## ADR-005: HTTP Framework Choice (Huma + Chi)

**Status**: Accepted

**Context**: Need lightweight, composable HTTP stack for financial API.

**Decision**: Use:
- **Chi**: Lightweight router with middleware support
- **Huma**: REST framework with automatic OpenAPI documentation

**Rationale**:
- Chi: Minimal, idiomatic Go, great middleware ecosystem
- Huma: Automatic OpenAPI/Swagger docs, type-safe handlers
- Both are lightweight and maintainable
- Good community support

**Consequences**:
- Limited to HTTP transport (not gRPC, etc.)
- Huma's magic may be opaque to newcomers
- But excellent for rapid API development

---

## Future ADRs

These decisions pending as requirements evolve:
- **Persistence**: Database technology and ORM choice
- **Testing**: Unit test framework and test patterns
- **Logging**: Structured logging approach
- **Deployment**: Containerization and orchestration
- **Auth**: Authentication and authorization strategy
