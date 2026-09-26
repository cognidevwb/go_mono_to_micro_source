package payments

import "github.com/acme/shop/internal/platform"

// Payment records one charge. Status is a bare string (legacy smell).
type Payment struct {
	ID      uint `gorm:"primaryKey"`
	OrderID uint `gorm:"index"`
	Amount  float64
	Status  string
}

func init() { platform.Register(&Payment{}) }
