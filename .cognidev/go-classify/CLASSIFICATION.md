# Go classification

Bank revision `2026-09-25.1` · model `jev-1.13.0` · Jev: `answered` · source consent: `yes`

- 59 units in 18 Go files, every one listed under **All units** below.
- 36 units have at least one category settled by a fact (listed under **By category**, source `fact`).
- 2 excluded entries (2 declarations), listed under **Excluded**.
- 12 units need review, listed under **Needs review**.
- 4 units have a kind code left open and Jev did not settle, listed under **Kind open**.
- Requests: 70 sent, 0 from cache.

## By category

### authentication (1)

- `internal/platform/auth.go`
  - `AuthMiddleware` (middleware, cross-cutting) — internal/platform/auth.go:17 — fact: stores the caller's identity on the request context: writes user key platform.UserIDKey via context.WithValue (ambient-context, line 24)

### authorization (0)

_none_

### business_rules (6)

- `internal/catalog/service.go`
  - `Service.Create` (service, application) — internal/catalog/service.go:30 — jev p=0.82
- `internal/customers/service.go`
  - `Service.IsActive` (service, application) — internal/customers/service.go:24 — jev p=0.76
- `internal/inventory/service.go`
  - `Service.Reserve` (service, application) — internal/inventory/service.go:18 — jev p=0.97
  - `Service.Restock` (service, application) — internal/inventory/service.go:36 — jev p=0.91
- `internal/orders/service.go`
  - `Service.PlaceOrder` (service, application) — internal/orders/service.go:36 — jev p=0.98
- `internal/payments/service.go`
  - `Service.Charge` (service, application) — internal/payments/service.go:15 — jev p=0.79

### caching (1)

- `internal/catalog/service.go`
  - `Service.PriceOf` (service, application) — internal/catalog/service.go:38 — fact: read cache catalog.priceCache (priceCache), line 40

### concurrency (6)

- `internal/catalog/service.go`
  - `Service.PriceOf` (service, application) — internal/catalog/service.go:38 — fact: calls .Lock(), .RLock(), .Unlock(), .RUnlock()
- `internal/inventory/service.go`
  - `StartRestockJob` (job, application) — internal/inventory/service.go:42 — fact: 1 go statement(s)
- `internal/platform/bus.go`
  - `Bus` (model, domain) — internal/platform/bus.go:10 — fact: uses type sync.RWMutex
  - `Bus.Subscribe` (open: service|util, unresolved: kind unresolved) — internal/platform/bus.go:19 — fact: calls .Lock(), .Unlock()
  - `Bus.Publish` (open: service|util, unresolved: kind unresolved) — internal/platform/bus.go:26 — fact: calls .RLock(), .RUnlock()
- `internal/platform/db.go`
  - `Register` (open: service|util, unresolved: kind unresolved) — internal/platform/db.go:12 — jev p=0.80

### configuration (1)

- `cmd/shop/main.go`
  - `main` (wiring, infrastructure) — cmd/shop/main.go:18 — fact: references os.Getenv

### error_mapping (1)

- `internal/customers/handlers.go`
  - `getCustomer` (handler, transport) — internal/customers/handlers.go:16 — jev p=0.78

### external_integration (2)

- `internal/payments/gateway.go`
  - `Gateway.Charge` (client, infrastructure) — internal/payments/gateway.go:23 — fact: Post /v1/charges via g.client, line 25
- `internal/payments/service.go`
  - `Service.Charge` (service, application) — internal/payments/service.go:15 — jev p=0.84

### file_io (0)

_none_

### http_transport (10)

- `internal/catalog/handlers.go`
  - `RegisterRoutes` (wiring, infrastructure) — internal/catalog/handlers.go:11 — fact: registers GET /api/products (gin)
  - `listProducts` (handler, transport) — internal/catalog/handlers.go:16 — fact: takes *github.com/gin-gonic/gin.Context
  - `createProduct` (handler, transport) — internal/catalog/handlers.go:25 — fact: takes *github.com/gin-gonic/gin.Context
- `internal/customers/handlers.go`
  - `RegisterRoutes` (wiring, infrastructure) — internal/customers/handlers.go:11 — fact: registers POST /api/customers (gin)
  - `getCustomer` (handler, transport) — internal/customers/handlers.go:16 — fact: takes *github.com/gin-gonic/gin.Context
  - `registerCustomer` (handler, transport) — internal/customers/handlers.go:30 — fact: takes *github.com/gin-gonic/gin.Context
- `internal/orders/handlers.go`
  - `RegisterRoutes` (wiring, infrastructure) — internal/orders/handlers.go:13 — fact: registers POST /api/orders (gin)
  - `placeOrder` (handler, transport) — internal/orders/handlers.go:18 — fact: takes *github.com/gin-gonic/gin.Context
  - `myOrders` (handler, transport) — internal/orders/handlers.go:33 — fact: takes *github.com/gin-gonic/gin.Context
- `internal/platform/auth.go`
  - `AuthMiddleware` (middleware, cross-cutting) — internal/platform/auth.go:17 — fact: returns github.com/gin-gonic/gin.HandlerFunc

### input_validation (7)

- `internal/catalog/model.go`
  - `Product` (model, domain) — internal/catalog/model.go:13 — fact: struct tags `binding`/`validate` declare the rules
- `internal/catalog/service.go`
  - `Service.Create` (service, application) — internal/catalog/service.go:30 — jev p=0.83
- `internal/customers/handlers.go`
  - `getCustomer` (handler, transport) — internal/customers/handlers.go:16 — jev p=0.75
- `internal/customers/model.go`
  - `Customer` (model, domain) — internal/customers/model.go:10 — fact: struct tags `binding`/`validate` declare the rules
- `internal/orders/model.go`
  - `PlaceOrderRequest` (dto, transport) — internal/orders/model.go:24 — fact: struct tags `binding`/`validate` declare the rules
  - `LineItem` (dto, transport) — internal/orders/model.go:30 — fact: struct tags `binding`/`validate` declare the rules
- `internal/platform/auth.go`
  - `AuthMiddleware` (middleware, cross-cutting) — internal/platform/auth.go:17 — jev p=0.72

### messaging (5)

- `cmd/shop/main.go`
  - `main` (wiring, infrastructure) — cmd/shop/main.go:18 — fact: subscribes inventory.Service.OnOrderPlaced to order.placed
- `internal/inventory/service.go`
  - `Service.OnOrderPlaced` (service, application) — internal/inventory/service.go:31 — fact: handles event order.placed (subscribed at cmd/shop/main.go:32)
- `internal/orders/service.go`
  - `Service.PlaceOrder` (service, application) — internal/orders/service.go:36 — fact: publishes order.placed via s.bus.Publish, line 73
- `internal/platform/bus.go`
  - `Bus.Subscribe` (open: service|util, unresolved: kind unresolved) — internal/platform/bus.go:19 — jev p=0.94
  - `Bus.Publish` (open: service|util, unresolved: kind unresolved) — internal/platform/bus.go:26 — jev p=0.95

### observability (0)

_none_

### persistence (19)

- `internal/catalog/model.go`
  - `Category` (model, domain) — internal/catalog/model.go:6 — fact: persisted entity (table:categories (convention))
  - `Product` (model, domain) — internal/catalog/model.go:13 — fact: persisted entity (table:products (convention))
- `internal/catalog/service.go`
  - `Service.List` (service, application) — internal/catalog/service.go:23 — fact: reads Product via s.db (direct) (entity-access, line 25)
  - `Service.Create` (service, application) — internal/catalog/service.go:30 — fact: writes Product via s.db (direct) (entity-access, line 34)
  - `Service.PriceOf` (service, application) — internal/catalog/service.go:38 — fact: reads Product via s.db (direct) (entity-access, line 46)
- `internal/customers/model.go`
  - `Customer` (model, domain) — internal/customers/model.go:10 — fact: persisted entity (table:customers (convention))
- `internal/customers/service.go`
  - `Service.Get` (service, application) — internal/customers/service.go:12 — fact: reads Customer via s.db (direct) (entity-access, line 14)
  - `Service.Register` (service, application) — internal/customers/service.go:21 — fact: writes Customer via s.db (direct) (entity-access, line 21)
  - `Service.IsActive` (service, application) — internal/customers/service.go:24 — jev p=0.78
- `internal/inventory/model.go`
  - `StockItem` (model, domain) — internal/inventory/model.go:6 — fact: persisted entity (table:stock_items)
- `internal/inventory/service.go`
  - `Service.Reserve` (service, application) — internal/inventory/service.go:18 — fact: reads StockItem via tx (direct) (entity-access, line 20)
  - `Service.Restock` (service, application) — internal/inventory/service.go:36 — fact: writes StockItem via s.db (direct) (entity-access, line 38)
- `internal/orders/model.go`
  - `Order` (model, domain) — internal/orders/model.go:6 — fact: persisted entity (table:orders (convention))
  - `OrderLine` (model, domain) — internal/orders/model.go:15 — fact: persisted entity (table:order_lines (convention))
- `internal/orders/service.go`
  - `Service.PlaceOrder` (service, application) — internal/orders/service.go:36 — fact: writes Order via tx (direct) (entity-access, line 61)
  - `Service.ForCustomer` (service, application) — internal/orders/service.go:78 — fact: reads Order via s.db (direct) (entity-access, line 80)
- `internal/payments/model.go`
  - `Payment` (model, domain) — internal/payments/model.go:6 — fact: persisted entity (table:payments (convention))
- `internal/payments/service.go`
  - `Service.Charge` (service, application) — internal/payments/service.go:15 — fact: writes Payment via tx (direct) (entity-access, line 19)
- `internal/platform/db.go`
  - `Open` (open: service|util, unresolved: kind unresolved) — internal/platform/db.go:15 — jev p=0.96

### scheduling (1)

- `internal/inventory/service.go`
  - `StartRestockJob` (job, application) — internal/inventory/service.go:42 — fact: references time.NewTicker

### secrets_crypto (0)

_none_

### serialization (1)

- `internal/payments/gateway.go`
  - `Gateway.Charge` (client, infrastructure) — internal/payments/gateway.go:23 — fact: references encoding/json.Marshal

### transactions (1)

- `internal/orders/service.go`
  - `Service.PlaceOrder` (service, application) — internal/orders/service.go:36 — fact: calls .Transaction()

## Needs review

Unresolved answers never feed a derivation; these units are listed by name in every playbook that routes them.

- `Category` (model, domain) — internal/catalog/model.go:6 — unresolved: serialization
- `registerCustomer` (handler, transport) — internal/customers/handlers.go:30 — unresolved: error_mapping
- `StockItem` (model, domain) — internal/inventory/model.go:6 — unresolved: serialization
- `Service.Restock` (service, application) — internal/inventory/service.go:36 — unresolved: scheduling
- `StartRestockJob` (job, application) — internal/inventory/service.go:42 — unresolved: persistence
- `Service.PlaceOrder` (service, application) — internal/orders/service.go:36 — unresolved: authorization, external_integration
- `AuthMiddleware` (middleware, cross-cutting) — internal/platform/auth.go:17 — unresolved: error_mapping
- `NewBus` (util, cross-cutting) — internal/platform/bus.go:16 — unresolved: messaging
- `Bus.Subscribe` (open: service|util, unresolved: kind unresolved) — internal/platform/bus.go:19 — unresolved: kind
- `Bus.Publish` (open: service|util, unresolved: kind unresolved) — internal/platform/bus.go:26 — unresolved: kind
- `Register` (open: service|util, unresolved: kind unresolved) — internal/platform/db.go:12 — unresolved: kind
- `Open` (open: service|util, unresolved: kind unresolved) — internal/platform/db.go:15 — unresolved: kind

## Kind open

- `Bus.Subscribe` (open: service|util, unresolved: kind unresolved) — internal/platform/bus.go:19 — unresolved (no fact or name settles it)
- `Bus.Publish` (open: service|util, unresolved: kind unresolved) — internal/platform/bus.go:26 — unresolved (no fact or name settles it)
- `Register` (open: service|util, unresolved: kind unresolved) — internal/platform/db.go:12 — unresolved (no fact or name settles it)
- `Open` (open: service|util, unresolved: kind unresolved) — internal/platform/db.go:15 — unresolved (no fact or name settles it)

## All units

- `cmd/shop/main.go`
  - `main` (wiring, infrastructure) — cmd/shop/main.go:18 — configuration, messaging, risk 0
- `internal/catalog/handlers.go`
  - `RegisterRoutes` (wiring, infrastructure) — internal/catalog/handlers.go:11 — http_transport, risk 0
  - `listProducts` (handler, transport) — internal/catalog/handlers.go:16 — http_transport, risk 2
  - `createProduct` (handler, transport) — internal/catalog/handlers.go:25 — http_transport, risk 2
- `internal/catalog/model.go`
  - `Category` (model, domain) — internal/catalog/model.go:6 — persistence
  - `Product` (model, domain) — internal/catalog/model.go:13 — input_validation, persistence
  - `init` (wiring, infrastructure) — internal/catalog/model.go:22 — no category, risk 0
- `internal/catalog/service.go`
  - `Service` (service, application) — internal/catalog/service.go:17 — no category
  - `NewService` (service, application) — internal/catalog/service.go:20 — no category, risk 0
  - `Service.List` (service, application) — internal/catalog/service.go:23 — persistence, risk 0
  - `Service.Create` (service, application) — internal/catalog/service.go:30 — persistence, input_validation, business_rules, risk 1
  - `Service.PriceOf` (service, application) — internal/catalog/service.go:38 — concurrency, persistence, caching, risk 0
- `internal/customers/handlers.go`
  - `RegisterRoutes` (wiring, infrastructure) — internal/customers/handlers.go:11 — http_transport, risk 0
  - `getCustomer` (handler, transport) — internal/customers/handlers.go:16 — http_transport, input_validation, error_mapping, risk 2
  - `registerCustomer` (handler, transport) — internal/customers/handlers.go:30 — http_transport, risk 2
- `internal/customers/model.go`
  - `Customer` (model, domain) — internal/customers/model.go:10 — input_validation, persistence
  - `init` (wiring, infrastructure) — internal/customers/model.go:18 — no category, risk 0
- `internal/customers/service.go`
  - `Service` (service, application) — internal/customers/service.go:6 — no category
  - `NewService` (service, application) — internal/customers/service.go:9 — no category, risk 0
  - `Service.Get` (service, application) — internal/customers/service.go:12 — persistence, risk 1
  - `Service.Register` (service, application) — internal/customers/service.go:21 — persistence, risk 1
  - `Service.IsActive` (service, application) — internal/customers/service.go:24 — persistence, business_rules, risk 0
- `internal/inventory/model.go`
  - `StockItem` (model, domain) — internal/inventory/model.go:6 — persistence
  - `StockItem.TableName` (model, domain) — internal/inventory/model.go:14 — no category, risk 0
  - `init` (wiring, infrastructure) — internal/inventory/model.go:16 — no category, risk 0
- `internal/inventory/service.go`
  - `Service` (service, application) — internal/inventory/service.go:12 — no category
  - `NewService` (service, application) — internal/inventory/service.go:15 — no category, risk 0
  - `Service.Reserve` (service, application) — internal/inventory/service.go:18 — persistence, business_rules, risk 1
  - `Service.OnOrderPlaced` (service, application) — internal/inventory/service.go:31 — messaging, risk 0
  - `Service.Restock` (service, application) — internal/inventory/service.go:36 — persistence, business_rules, risk 1
  - `StartRestockJob` (job, application) — internal/inventory/service.go:42 — scheduling, concurrency, risk 1
- `internal/orders/handlers.go`
  - `RegisterRoutes` (wiring, infrastructure) — internal/orders/handlers.go:13 — http_transport, risk 0
  - `placeOrder` (handler, transport) — internal/orders/handlers.go:18 — http_transport, risk 2
  - `myOrders` (handler, transport) — internal/orders/handlers.go:33 — http_transport, risk 2
- `internal/orders/model.go`
  - `Order` (model, domain) — internal/orders/model.go:6 — persistence
  - `OrderLine` (model, domain) — internal/orders/model.go:15 — persistence
  - `PlaceOrderRequest` (dto, transport) — internal/orders/model.go:24 — input_validation
  - `LineItem` (dto, transport) — internal/orders/model.go:30 — input_validation
  - `init` (wiring, infrastructure) — internal/orders/model.go:35 — no category, risk 0
- `internal/orders/service.go`
  - `Service` (service, application) — internal/orders/service.go:19 — no category
  - `NewService` (service, application) — internal/orders/service.go:29 — no category, risk 0
  - `Service.PlaceOrder` (service, application) — internal/orders/service.go:36 — transactions, persistence, messaging, business_rules, risk 2
  - `Service.ForCustomer` (service, application) — internal/orders/service.go:78 — persistence, risk 0
- `internal/payments/gateway.go`
  - `Gateway` (client, infrastructure) — internal/payments/gateway.go:12 — no category
  - `NewGateway` (client, infrastructure) — internal/payments/gateway.go:18 — no category, risk 0
  - `Gateway.Charge` (client, infrastructure) — internal/payments/gateway.go:23 — serialization, external_integration, risk 1
- `internal/payments/model.go`
  - `Payment` (model, domain) — internal/payments/model.go:6 — persistence
  - `init` (wiring, infrastructure) — internal/payments/model.go:13 — no category, risk 0
- `internal/payments/service.go`
  - `Service` (service, application) — internal/payments/service.go:6 — no category
  - `NewService` (service, application) — internal/payments/service.go:12 — no category, risk 0
  - `Service.Charge` (service, application) — internal/payments/service.go:15 — persistence, external_integration, business_rules, risk 1
- `internal/platform/auth.go`
  - `AuthMiddleware` (middleware, cross-cutting) — internal/platform/auth.go:17 — http_transport, authentication, input_validation, risk 0
  - `UserID` (util, cross-cutting) — internal/platform/auth.go:32 — no category, risk 0
- `internal/platform/bus.go`
  - `Bus` (model, domain) — internal/platform/bus.go:10 — concurrency
  - `NewBus` (util, cross-cutting) — internal/platform/bus.go:16 — no category, risk 0
  - `Bus.Subscribe` (open: service|util, unresolved: kind unresolved) — internal/platform/bus.go:19 — concurrency, messaging, risk 0
  - `Bus.Publish` (open: service|util, unresolved: kind unresolved) — internal/platform/bus.go:26 — concurrency, messaging, risk 0
- `internal/platform/db.go`
  - `Register` (open: service|util, unresolved: kind unresolved) — internal/platform/db.go:12 — concurrency, risk 0
  - `Open` (open: service|util, unresolved: kind unresolved) — internal/platform/db.go:15 — persistence, risk 0

## Excluded

### named non-struct type (not a unit: no fields, no behaviour) (2)

- `internal/platform/auth.go#ctxKey` — internal/platform/auth.go:11
- `internal/platform/bus.go#Handler` — internal/platform/bus.go:6

