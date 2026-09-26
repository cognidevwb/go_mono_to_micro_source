package customers

import (
	"time"

	"github.com/acme/shop/internal/platform"
)

// Customer places orders.
type Customer struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Email     string    `gorm:"uniqueIndex;size:200" json:"email" binding:"required,email"`
	Name      string    `json:"name" binding:"required"`
	Active    bool      `gorm:"default:true" json:"active"`
	CreatedAt time.Time `json:"createdAt"`
}

func init() { platform.Register(&Customer{}) }
