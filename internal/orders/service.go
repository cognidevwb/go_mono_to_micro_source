package orders

import (
	"errors"

	"gorm.io/gorm"

	"github.com/acme/shop/internal/catalog"
	"github.com/acme/shop/internal/customers"
	"github.com/acme/shop/internal/inventory"
	"github.com/acme/shop/internal/payments"
	"github.com/acme/shop/internal/platform"
)

// TopicOrderPlaced is published after an order commits.
const TopicOrderPlaced = "order.placed"

// Service owns the orders context and orchestrates the others.
type Service struct {
	db        *gorm.DB
	catalog   *catalog.Service
	customers *customers.Service
	inventory *inventory.Service
	payments  *payments.Service
	bus       *platform.Bus
}

// NewService wires the order service.
func NewService(db *gorm.DB, c *catalog.Service, cu *customers.Service,
	i *inventory.Service, p *payments.Service, bus *platform.Bus) *Service {
	return &Service{db: db, catalog: c, customers: cu, inventory: i, payments: p, bus: bus}
}

// PlaceOrder spans customers, inventory, catalog and payments in ONE
// transaction — the CreateOrder saga seam.
func (s *Service) PlaceOrder(req PlaceOrderRequest) (*Order, error) {
	active, err := s.customers.IsActive(req.CustomerID)
	if err != nil {
		return nil, err
	}
	if !active {
		return nil, errors.New("customer inactive")
	}
	order := &Order{CustomerID: req.CustomerID, Status: "pending"}
	err = s.db.Transaction(func(tx *gorm.DB) error {
		for _, item := range req.Items {
			ok, err := s.inventory.Reserve(tx, item.ProductID, item.Quantity)
			if err != nil {
				return err
			}
			if !ok {
				return errors.New("out of stock")
			}
			price, err := s.catalog.PriceOf(item.ProductID)
			if err != nil {
				return err
			}
			order.Lines = append(order.Lines, OrderLine{ProductID: item.ProductID, Quantity: item.Quantity, UnitPrice: price})
			order.Total += price * float64(item.Quantity)
		}
		if err := tx.Create(order).Error; err != nil {
			return err
		}
		if err := s.payments.Charge(tx, order.ID, order.Total); err != nil {
			return err
		}
		order.Status = "placed"
		return tx.Save(order).Error
	})
	if err != nil {
		return nil, err
	}
	s.bus.Publish(TopicOrderPlaced, order.ID)
	return order, nil
}

// ForCustomer lists a customer's orders.
func (s *Service) ForCustomer(customerID uint) ([]Order, error) {
	var out []Order
	err := s.db.Preload("Lines").Where("customer_id = ?", customerID).Find(&out).Error
	return out, err
}
