# abstract-mongodb

English — Quick overview

abstract-mongodb is a small Go helper library that provides a thin abstraction over the
official MongoDB Go driver (`go.mongodb.org/mongo-driver`). It centralizes connection
management and common CRUD patterns to make implementing repository code easier.

Key points
- Singleton connection via `GetDB()` (see `db.go`). Environment variables are required:
	- `MONGODB_URI` — MongoDB connection URI
	- `MONGODB_DB_NAME` — database name
	- If `MODE` is `DEBUG` or `TEST`, the library appends `_TEST` to the DB name.
- Main types are in `struct.go`: `MongoDB`, `FindParams`, `CountDocumentsParams`, `FunctionTransaction`.
- Helper functions:
	- `FindOne`, `Find`, `FindAggregate` in `find.go` (use `FindParams`).
	- `InsertOne` (`insert.go`), `UpdateOne` / `UpdateMany` (`update.go`), `CountDocuments` (`count.go`).
	- `ExecTransactionFunctions` (`transaction.go`) accepts `FunctionTransaction` slices and runs them inside `WithTransaction`.
	- Projection builders and index helpers are in `utils.go` (`BuildFindOptions`, `BuildFindOneOptions`, `EnsureIndex`).

Usage examples

Minimal `.env` example:

```env
MONGODB_URI=mongodb://localhost:27017
MONGODB_DB_NAME=myapp
MODE=DEBUG
```

Build

```bash
go build ./...
```

Example: finding a single document

```go
var result MyModel
err := mongodb.FindOne(ctx, &result, mongodb.FindParams{
		Collection: "users",
		Filter: bson.M{"email": "foo@example.com"},
		Fields: []string{"_id", "email", "name"},
})
```

Transactions (pattern)

```go
f1 := mongodb.FunctionTransaction{
		Function: func(sessCtx mongo.SessionContext, params ...any) error {
				// use sessCtx as context for DB ops
				return nil
		},
		Params: []any{},
}
_ = mongodb.ExecTransactionFunctions(ctx, f1)
```

Where to look
- `db.go` — connection singleton and `.env` behavior
- `struct.go` — core types
- `find.go`, `insert.go`, `update.go`, `count.go`, `transaction.go`, `utils.go` — main helpers

Português (PT-BR) — Visão rápida

`abstract-mongodb` é uma biblioteca auxiliar em Go que fornece uma camada leve sobre
o driver oficial do MongoDB (`go.mongodb.org/mongo-driver`). Centraliza o gerenciamento
de conexão e padrões CRUD comuns para facilitar a implementação de repositórios.

Pontos principais
- Conexão singleton através de `GetDB()` (veja `db.go`). Variáveis de ambiente necessárias:
	- `MONGODB_URI` — URI de conexão do MongoDB
	- `MONGODB_DB_NAME` — nome do banco de dados
	- Se `MODE` for `DEBUG` ou `TEST`, o nome do DB recebe sufixo `_TEST`.
- Tipos principais em `struct.go`: `MongoDB`, `FindParams`, `CountDocumentsParams`, `FunctionTransaction`.
- Funções de ajuda:
	- `FindOne`, `Find`, `FindAggregate` em `find.go` (usa `FindParams`).
	- `InsertOne` (`insert.go`), `UpdateOne` / `UpdateMany` (`update.go`), `CountDocuments` (`count.go`).
	- `ExecTransactionFunctions` (`transaction.go`) aceita um slice de `FunctionTransaction` e executa dentro de `WithTransaction`.
	- Builders de projeção e utilitários de índices em `utils.go` (`BuildFindOptions`, `BuildFindOneOptions`, `EnsureIndex`).

Exemplos de uso

Arquivo `.env` mínimo:

```env
MONGODB_URI=mongodb://localhost:27017
MONGODB_DB_NAME=myapp
MODE=DEBUG
```

Build

```bash
go build ./...
```

Exemplo: buscar um documento

```go
var result MyModel
err := mongodb.FindOne(ctx, &result, mongodb.FindParams{
		Collection: "users",
		Filter: bson.M{"email": "foo@example.com"},
		Fields: []string{"_id", "email", "name"},
})
```

Transações (padrão)

```go
f1 := mongodb.FunctionTransaction{
		Function: func(sessCtx mongo.SessionContext, params ...any) error {
				// use sessCtx como contexto para operações no DB
				return nil
		},
		Params: []any{},
}
_ = mongodb.ExecTransactionFunctions(ctx, f1)
```

Contribuindo

1. Follow existing patterns: always use `context.Context` and `GetDB()` for DB access.
2. When adding queries, use `FindParams` + `BuildFindOptions` to support projection consistency.
3. If you add tests that hit a database, set `MODE=TEST` so the DB name is suffixed with `_TEST`.

License

This project includes a `LICENSE` file in the repository root.
