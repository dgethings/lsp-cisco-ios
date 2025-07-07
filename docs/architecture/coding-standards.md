# Coding Standards

This document outlines the coding standards for the Cisco IOS LSP project. Adherence to these standards is mandatory for all development.

## Core Standards

- **Languages & Runtimes:** Go 1.22.1
- **Style & Linting:** Adhere to `gofmt` and `golint` recommendations. Use `slog` for logging.
- **Test Organization:** Tests should be placed in `_test.go` files within the same package as the code they test.

## Naming Conventions

- **Variables:** `camelCase`
- **Functions:** `CamelCase` for exported functions, `camelCase` for unexported functions.
- **Packages:** `lowercase`
- **Files:** `snake_case` for Go source files.

## Critical Rules

- All external inputs MUST be validated.
- NEVER hardcode secrets; use configuration.
- Ensure proper error handling for all operations.
- Avoid global state where possible; pass dependencies explicitly.
- All new code must have corresponding unit tests.

## Language-Specific Guidelines

### Go Specifics

- Use Go modules for dependency management.
- Prefer composition over inheritance.
- Handle errors explicitly; avoid `panic` for recoverable errors.
- Use `context.Context` for request-scoped values, cancellation, and deadlines.
