package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/gin-gonic/gin"

	"github.com/acme/shop/internal/catalog"
	"github.com/acme/shop/internal/customers"
	"github.com/acme/shop/internal/inventory"
	"github.com/acme/shop/internal/orders"
	"github.com/acme/shop/internal/payments"
	"github.com/acme/shop/internal/platform"
)

func main() {
	db, err := platform.Open(os.Getenv("DATABASE_URL"))
	if err != nil {
		slog.Error("open database", "err", err)
		os.Exit(1)
	}
	bus := platform.NewBus()

	catalogSvc := catalog.NewService(db)
	customerSvc := customers.NewService(db)
	inventorySvc := inventory.NewService(db)
	paymentSvc := payments.NewService(db, payments.NewGateway(os.Getenv("PAYMENTS_URL")))
	orderSvc := orders.NewService(db, catalogSvc, customerSvc, inventorySvc, paymentSvc, bus)

	bus.Subscribe(orders.TopicOrderPlaced, inventorySvc.OnOrderPlaced)
	inventory.StartRestockJob(context.Background(), inventorySvc)

	r := gin.Default()
	r.Use(platform.AuthMiddleware())
	api := r.Group("/api")
	catalog.RegisterRoutes(api, catalogSvc)
	customers.RegisterRoutes(api, customerSvc)
	orders.RegisterRoutes(api, orderSvc)

	if err := r.Run(":8080"); err != nil {
		slog.Error("server stopped", "err", err)
	}
}
