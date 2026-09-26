package payments

import "gorm.io/gorm"

// Service owns the payments context.
type Service struct {
	db      *gorm.DB
	gateway *Gateway
}

// NewService builds the payment service.
func NewService(db *gorm.DB, gw *Gateway) *Service { return &Service{db: db, gateway: gw} }

// Charge captures amount and records the payment inside tx.
func (s *Service) Charge(tx *gorm.DB, orderID uint, amount float64) error {
	if err := s.gateway.Charge(orderID, amount); err != nil {
		return err
	}
	return tx.Create(&Payment{OrderID: orderID, Amount: amount, Status: "captured"}).Error
}
