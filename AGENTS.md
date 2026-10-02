# alutec-inventory-be — usecase and handler structure

Canonical reference: `internal/usecases/product/create.go` and
`internal/usecases/stock_movement/create.go` (the one aggregate operation
that spans two tables).

This repo mirrors `practiq-be`'s architecture. If something here is
ambiguous, that repo is the tie-breaker.

## 1. One file per public use case

The file is named after the operation, not the domain:

```sh
internal/usecases/<domain>/
  create.go
  get.go
  list.go
  update.go
  delete.go
```

Never group operations in `manage.go`, `handlers.go`, `repository.go`, or any
file named after the domain instead of the operation. `stock_movement` only
has `create.go` and `list.go` — there is no `update.go` or `delete.go` for
it, and there must never be one (see §11).

## 2. What goes in an operation file

```go
type (
    CreateUsecase interface {
        Execute(ctx context.Context, in CreateInput) (*CreateOutput, apperrors.ApplicationError)
    }

    createUsecase struct {
        contextFactory appcontext.Factory
    }

    CreateInput struct {
        Name string `json:"name" binding:"required"`
    }

    CreateOutput struct {
        Data CategoryData `json:"data"`
    }
)

func NewCreateUsecase(contextFactory appcontext.Factory) CreateUsecase { ... }

func (u *createUsecase) Execute(...) (*CreateOutput, apperrors.ApplicationError) { ... }
```

- **One** interface per file, with a **single** public method: `Execute`.
- Private struct with the same name, lowercased.
- Constructor `New<Operation>Usecase`.
- `Input` and `Output` belong to the operation. Shared `Data` types live in
  `output_types.go`.

There is no authentication or multi-tenancy in this project (internal tool,
single workshop). `Execute` therefore takes only `ctx`, path ids, and the
`Input` — never a `requesterID` or `isSuperAdmin` like `practiq-be`. If this
app ever grows multi-user auth, add identity as leading parameters the same
way `practiq-be` does, don't smuggle it into the Input.

## 3. Helpers

- A helper used by **one** operation lives in that operation's file.
- A helper shared by several goes in a file with an explicit name
  (`output_types.go`, `validation.go`, etc).
- A helper file is fine; a file holding several operations is not.

## 4. What a refactor must not change

Moving code is not changing behaviour. These stay identical:

- HTTP routes and their verbs.
- The shape of request and response JSON, field by field.
- HTTP status codes and internal error codes.
- Validation, transactions, and the order of operations.

If none of that changes, **do not touch the frontend**
(`../alutec-inventory-fe`).

## 5. After each domain

```bash
gofmt -w ./internal/...
go vet ./...
go test ./internal/usecases/<domain>/...
go build ./...
```

## 6. Comments

This repo carries no explanatory comments in the code, except where a
non-obvious invariant needs one (e.g. "append-only ledger, no update/delete
by design"). The reasoning for a change goes in the commit message.

## 7. Errors

Error details are declared in
`internal/platform/errors/mappings/domain.go`, never inline in a usecase:

```go
// mappings/domain.go
ProductCreateError = ErrorDetails{
    InternalCode: "product:create:error",
    StatusCode:   http.StatusInternalServerError,
    Message:      "failed to create product",
}

// usecase
return nil, apperrors.NewApplicationError(mappings.ProductCreateError, err)
```

Use `http.StatusX` constants, not bare numbers. For unique/foreign-key
constraint violations, check them explicitly via `platform/pgerr` and return
a `apperrors.NewConflictError(...)` or `apperrors.NewBadRequestError(...)`
instead of the generic 500 mapping — see `usecases/product/create.go`.

## 8. Inputs live in the usecase

The `Input` is the use case's contract and lives **in the usecase**, with
`json` tags. The handler binds straight into it.

```go
// handlers/product/create.go
var input ucProduct.CreateInput
if err := c.ShouldBindJSON(&input); err != nil { ... }

output, appErr := uc.Execute(c, input)
```

Rules:

- **The Input is the body and nothing else.** Nothing from a path or query
  param belongs inside it — `list.go` usecases take a dedicated `ListFilter`
  struct built by the handler from query params instead (see
  `usecases/product/list.go`).
- **Path ids go as parameters**, before the Input: `Execute(ctx, id, input)`.
- **There is no `CreateCommand` or `CreateParams`.** `json` tags do not tie an
  Input to HTTP.
- **No `input_types.go` in handlers.** Duplicating the struct to copy it
  field by field is work nobody reads.

## 9. HTTP handlers

One endpoint per file, named after the operation:

```sh
internal/adapters/web/handlers/<domain>/
  create.go
  get.go
  list.go
  update.go
  delete.go
```

- Each file exports only `New<Operation>Handler`.
- The handler does one job: bind the body (or query params) into the
  usecase's Input/filter, call `Execute`, and respond.
- Business logic and validation live in the usecase.
- Never create `handlers.go`, `handler.go`, or files holding several
  endpoints.

A handler refactor must not change routes, verbs, JSON, status codes or
error codes.

## 10. Repositories

Canonical reference: `internal/adapters/datasources/repositories/product/`.

One package per table, one file per query, named after the operation:

```sh
internal/adapters/datasources/repositories/<domain>/
  repository.go        interface, struct, constructor, filter options
  create.go
  get.go
  list.go
  update.go
  delete.go
```

`repository.go` holds only the seams, never SQL:

```go
type ListFilterOptions struct {
    CategoryID string
    SupplierID string
    Search     string
    LowStock   bool
}

type Repository interface {
    Create(context.Context, domain.Product) (string, error)
    Get(context.Context, string) (*domain.Product, error)
    List(context.Context, ListFilterOptions) ([]domain.Product, error)
}

type repository struct {
    db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
    return &repository{db: db}
}
```

### Domain types

An entity read from or written to the database belongs in `internal/domain`,
never in a repository package. The repository imports that domain type in
its interface and operation files. `json` tags do not belong on domain
entities.

Rules:

- **SQL lives here and nowhere else.** A usecase never builds a query.
- **Return domain types**, never rows or driver types. A row that does not
  exist is `nil, nil`, not an error.
- **Return the raw `error`.** `apperrors.ApplicationError` belongs to the
  usecase.
- **Filters are named structs** (`ListFilterOptions`), not positional
  arguments.
- **Context is always the first argument**, through `QueryRowContext` /
  `QueryContext` / `ExecContext`, never `Query` or `Exec`.
- A new repository is registered in `repositories.go`, the only file that
  knows the whole set.

Small repositories with two or three methods may keep them in
`repository.go`; `category`, `supplier`, and `warehouse` do. Once a package
grows past that (`product`, `stock_movement`), split it by operation.

## 11. `stock_movement` is an append-only ledger — read this before touching it

`stock_movements` records what happened to stock and why. It is **never**
updated or deleted by the application:

- There is no `Update` or `Delete` method on
  `repositories/stock_movement.Repository`, no `UpdateUsecase` /
  `DeleteUsecase` in `usecases/stock_movement`, and no `PUT`/`DELETE` route
  for it in `routes.go`. Do not add any of the three. If a mistake needs
  correcting, the fix is creating a new movement that reverses it — not
  mutating history.
- `stock_movement.Create` is the **only** place that changes
  `products.stock_qty`, and it does so inside a single SQL transaction
  together with the `INSERT` (`repositories/stock_movement/create.go`). Any
  future change to how stock is adjusted must stay inside that transaction —
  never adjust stock from a usecase or from a different repository call, or
  the ledger and the balance can drift apart.
- `type = 'out'` subtracts `quantity`; `'in'` and `'adjustment'` add it. This
  matches the frontend's logic in `alutec-inventory-fe` exactly — keep both
  sides in sync if it ever changes.

## 12. Database

- Raw SQL via `database/sql` + `lib/pq`, no ORM. Migrations via
  `golang-migrate`, files in `migrations/000NNN_name.{up,down}.sql`.
- Every table's primary key is `UUID DEFAULT uuid_generate_v4()` (the
  `uuid-ossp` extension is created once in `000001_init_schema.up.sql`).
- Unique/foreign-key constraints get an explicit name
  (`idx_products_sku`, etc.) so `platform/pgerr.IsUniqueViolation` can match
  on it from the usecase — never rely on Postgres's auto-generated name.

## 13. Running this project

```bash
cp .env.example .env
docker-compose up -d alutec-inventory-postgres-db
make run            # or: make run-dev (air)
```

See `README.md` for the full command list and the data model summary.
