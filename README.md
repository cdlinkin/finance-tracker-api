# Finance Tracker API

A REST API for tracking income and expenses. Built with Go as a pet project to learn backend development and clean architecture.

> Currently uses in-memory storage. PostgreSQL support is in progress.

## Why I built this

I wanted to practice structuring a real Go project — not just writing functions in a single file, but separating concerns the way it's done in production codebases.

## Project structure

```
cmd/api/
└── main.go

internal/
├── domain/        # types shared across layers
├── handler/       # HTTP: decode request, call service, encode response
├── service/       # business logic and validation
└── repository/    # data access (in-memory for now, PostgreSQL coming)
```

Each layer only knows about the one below it. The handler doesn't know where data is stored. The repository doesn't know anything about HTTP. This made it easy to reason about the code while building it.

## Stack

- Go 1.25.2
- [chi](https://github.com/go-chi/chi) — router
- in-memory storage (map) → PostgreSQL (in progress)

## Run locally

```bash
git clone https://github.com/cdlinkin/finance-tracker-api
cd finance-tracker-api
go run cmd/api/main.go
```

Server starts on `:3000`.

## Endpoints

| Method | Path | Description |
|--------|------|-------------|
| POST | /transaction | Add income or expense |
| GET | /transaction | List all transactions |
| GET | /transaction/summary | Total income, expense, balance |
| DELETE | /transaction/{id} | Delete a transaction |

### Example

```bash
curl -X POST http://localhost:3000/transaction \
  -H "Content-Type: application/json" \
  -d '{"type":"income","amount":50000,"category":"salary"}'

# {"id":1,"type":"income","amount":50000,"category":"salary","created_at":"..."}
```

## What's next

- [ ] PostgreSQL persistence
- [ ] Docker + docker-compose
- [ ] Unit tests
- [ ] Swagger docs