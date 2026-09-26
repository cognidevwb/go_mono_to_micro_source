package customers

import "gorm.io/gorm"

// Service owns the customers context.
type Service struct{ db *gorm.DB }

// NewService builds the customer service.
func NewService(db *gorm.DB) *Service { return &Service{db: db} }

// Get loads one customer.
func (s *Service) Get(id uint) (*Customer, error) {
	var c Customer
	if err := s.db.First(&c, id).Error; err != nil {
		return nil, err
	}
	return &c, nil
}

// Register stores a new customer.
func (s *Service) Register(c *Customer) error { return s.db.Create(c).Error }

// IsActive reports whether the customer may order.
func (s *Service) IsActive(id uint) (bool, error) {
	c, err := s.Get(id)
	if err != nil {
		return false, err
	}
	return c.Active, nil
}
