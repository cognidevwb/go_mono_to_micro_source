package inventory

import "github.com/acme/shop/internal/platform"

// StockItem is on-hand stock for one product.
type StockItem struct {
	ID        uint `gorm:"primaryKey"`
	ProductID uint `gorm:"uniqueIndex"`
	OnHand    int
	Reserved  int
}

// TableName pins the legacy table name.
func (StockItem) TableName() string { return "stock_items" }

func init() { platform.Register(&StockItem{}) }
