package catalog

import (
	"errors"
	"sync"

	"gorm.io/gorm"
)

// priceCache is process-local shared state: two service instances disagree.
var (
	priceMu    sync.RWMutex
	priceCache = map[uint]float64{}
)

// Service owns the catalog context.
type Service struct{ db *gorm.DB }

// NewService builds the catalog service.
func NewService(db *gorm.DB) *Service { return &Service{db: db} }

// List returns every product.
func (s *Service) List() ([]Product, error) {
	var products []Product
	err := s.db.Find(&products).Error
	return products, err
}

// Create stores a new product.
func (s *Service) Create(p *Product) error {
	if p.Price <= 0 {
		return errors.New("price must be positive")
	}
	return s.db.Create(p).Error
}

// PriceOf returns a product's price, cached in-process.
func (s *Service) PriceOf(productID uint) (float64, error) {
	priceMu.RLock()
	price, ok := priceCache[productID]
	priceMu.RUnlock()
	if ok {
		return price, nil
	}
	var p Product
	if err := s.db.First(&p, productID).Error; err != nil {
		return 0, err
	}
	priceMu.Lock()
	priceCache[productID] = p.Price
	priceMu.Unlock()
	return p.Price, nil
}
