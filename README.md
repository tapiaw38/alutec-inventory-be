# alutec-inventory-be

Backend de inventario de Alutec Carpintería — Go con arquitectura limpia
(dominio / casos de uso / repositorios), sin ORM, SQL crudo vía `database/sql`.

Frontend: `../alutec-inventory-fe` (Quasar/Vue). Ver `AGENTS.md` para las
convenciones de este repo antes de tocar código.

## Requisitos

- Go 1.24+
- PostgreSQL 14+
- [golang-migrate CLI](https://github.com/golang-migrate/migrate)

## Configuración

```bash
cp .env.example .env
```

| Variable       | Descripción                      | Default |
|----------------|-----------------------------------|---------|
| `PORT`         | Puerto HTTP                       | `8080`  |
| `DATABASE_URL` | DSN de PostgreSQL                 | ver `.env.example` |
| `FRONTEND_URL` | Origen permitido por CORS         | `http://localhost:9000` |

## Base de datos

```bash
docker-compose up -d alutec-inventory-postgres-db
```

## Migraciones

```bash
make migrate-up      # aplicar todas
make migrate-down    # revertir la última
make migrate-create   # crear una nueva
```

La primera migración (`000001_init_schema`) crea `categories`, `suppliers`,
`warehouses`, `products` y `stock_movements`.

## Ejecución

```bash
make run        # go run
make run-dev    # con air (live reload)
```

El servidor también corre las migraciones pendientes al arrancar
(`database.RunMigrations`), así que `make migrate-up` es solo necesario para
gestión manual.

## Modelo de datos

- **categories**: nombre único + tipo (`raw_material` | `finished_good`).
- **suppliers**: proveedores.
- **warehouses**: depósitos.
- **products**: SKU único (también usado como valor de código de barras
  CODE128 en el frontend), categoría y proveedor por FK, stock actual.
- **stock_movements**: **ledger append-only**. No hay endpoint de edición ni
  borrado a propósito — crear un movimiento (`in` | `out` | `adjustment`) es
  la única forma de cambiar `products.stock_qty`, y lo hace dentro de una
  transacción (ver `internal/adapters/datasources/repositories/stock_movement/create.go`).
  Para corregir un error se carga un movimiento que lo revierte, nunca se
  edita el historial.

## API

Rutas base `/api`: `categories`, `suppliers`, `warehouses`, `products`
(CRUD completo), `stock-movements` (solo `POST` y `GET`, sin `PUT`/`DELETE`
por diseño).

`GET /health` para chequeo de liveness.

Sin autenticación en esta primera versión (herramienta interna). Si se agrega
multi-usuario más adelante, seguir el patrón de `appcontext.Factory` para
inyectar identidad, igual que en `practiq-be`.
