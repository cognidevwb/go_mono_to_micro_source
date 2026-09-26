# go-monolith-shop — Go (gin + gorm) monolith fixture

The Go twin of `dotnet-monolith-shop`: a small but genuine gin + gorm monolith
with real bounded-context seams, for exercising the Go Understand producers and
**modernization-go-monolith** end to end.

- Packages = candidate contexts: `catalog`, `customers`, `inventory`,
  `payments`, `orders` under `internal/`.
- One shared `*gorm.DB` (`platform.Open`) migrates every model — the
  database-per-service seam.
- `orders.Service.PlaceOrder` reaches into catalog + customers + inventory +
  payments and commits in **one `db.Transaction`** — the CreateOrder saga seam.
- Hidden coupling a split breaks without a compiler error:
  - `catalog.priceCache` — a package-level `map` guarded by a `sync.RWMutex`
    (shared mutable state);
  - `platform.Bus` — an in-process event bus (`OrderPlaced` published by orders,
    consumed by inventory);
  - `platform.UserIDKey` — a request-scoped `context.Context` value set by the
    auth middleware and read by orders;
  - `payments.Gateway` — an outbound HTTP call to a payment provider;
  - `inventory.StartRestockJob` — a background ticker goroutine.
- Deliberate legacy smells: `inventory.Service.Reserve` reads, checks, then
  writes (oversell); `payments.Payment.Status` is a bare string.

Decomposes to: catalog / customers / inventory / payments / orders services,
database-per-service, and the CreateOrder saga (reserve stock → charge payment).

Build: `go mod tidy && go build ./...`
