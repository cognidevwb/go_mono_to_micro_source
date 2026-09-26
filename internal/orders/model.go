package orders

import "github.com/acme/shop/internal/platform"

// Order is a customer's purchase.
type Order struct {
	ID         uint        `gorm:"primaryKey" json:"id"`
	CustomerID uint        `gorm:"index" json:"customerId"`
	Status     string      `json:"status"`
	Total      float64     `json:"total"`
	Lines      []OrderLine `gorm:"foreignKey:OrderID" json:"lines"`
}

// OrderLine is one product on an order.
type OrderLine struct {
	ID        uint    `gorm:"primaryKey" json:"id"`
	OrderID   uint    `json:"orderId"`
	ProductID uint    `json:"productId"`
	Quantity  int     `json:"quantity"`
	UnitPrice float64 `json:"unitPrice"`
}

// PlaceOrderRequest is the POST /orders body.
type PlaceOrderRequest struct {
	CustomerID uint       `json:"customerId" binding:"required"`
	Items      []LineItem `json:"items" binding:"required,min=1,dive"`
}

// LineItem is one requested product.
type LineItem struct {
	ProductID uint `json:"productId" binding:"required"`
	Quantity  int  `json:"quantity" binding:"required,gt=0"`
}

func init() { platform.Register(&Order{}, &OrderLine{}) }
