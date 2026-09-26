package catalog

import "github.com/acme/shop/internal/platform"

// Category groups products.
type Category struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	Name     string `gorm:"size:120;not null" json:"name"`
	Products []Product `json:"products,omitempty"`
}

// Product is a sellable item.
type Product struct {
	ID         uint    `gorm:"primaryKey" json:"id"`
	SKU        string  `gorm:"uniqueIndex;size:40" json:"sku" binding:"required"`
	Name       string  `gorm:"size:200" json:"name" binding:"required"`
	Price      float64 `json:"price" binding:"required,gt=0"`
	CategoryID uint    `json:"categoryId"`
	Category   Category `gorm:"foreignKey:CategoryID" json:"-"`
}

func init() { platform.Register(&Category{}, &Product{}) }
