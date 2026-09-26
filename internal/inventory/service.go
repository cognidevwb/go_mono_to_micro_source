package inventory

import (
	"context"
	"log"
	"time"

	"gorm.io/gorm"
)

// Service owns the inventory context.
type Service struct{ db *gorm.DB }

// NewService builds the inventory service.
func NewService(db *gorm.DB) *Service { return &Service{db: db} }

// Reserve holds qty units. Read-check-then-write: two orders can oversell.
func (s *Service) Reserve(tx *gorm.DB, productID uint, qty int) (bool, error) {
	var item StockItem
	if err := tx.Where("product_id = ?", productID).First(&item).Error; err != nil {
		return false, err
	}
	if item.OnHand-item.Reserved < qty {
		return false, nil
	}
	item.Reserved += qty
	return true, tx.Save(&item).Error
}

// OnOrderPlaced reacts to the orders context's OrderPlaced event.
func (s *Service) OnOrderPlaced(payload any) {
	log.Printf("order placed: %v", payload)
}

// Restock tops up every item below its reserve.
func (s *Service) Restock() error {
	return s.db.Model(&StockItem{}).Where("on_hand < reserved").
		Update("on_hand", gorm.Expr("reserved + 10")).Error
}

// StartRestockJob runs Restock every minute until ctx ends.
func StartRestockJob(ctx context.Context, s *Service) {
	go func() {
		t := time.NewTicker(time.Minute)
		defer t.Stop()
		for {
			select {
			case <-ctx.Done():
				return
			case <-t.C:
				if err := s.Restock(); err != nil {
					log.Printf("restock failed: %v", err)
				}
			}
		}
	}()
}
