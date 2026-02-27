# markdown-notes-api

A REST API for managing markdown notes.

## Requirements

- Go 1.22+

## Run locally

```bash
go run ./cmd/api
```

## Build

```bash
go build ./cmd/api
```

## Project structure

- `cmd/api/` - application entrypoint
- `internal/config/` - configuration loading
- `internal/handler/` - HTTP handlers
- `internal/middleware/` - middleware
- `internal/model/` - domain models
- `internal/routes/` - route registration
- `internal/service/` - business logic
- `internal/util/` - shared utilities
