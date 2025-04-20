package repository

import (
	"airline-management-system/internal/models/payment"

	"gorm.io/gorm"
)

type PaymentRepository struct {
	db *gorm.DB
}

func NewPaymentRepository(db *gorm.DB) *PaymentRepository {
	return &PaymentRepository{db}
}

func (r *PaymentRepository) GetAllPayments() ([]payment.Payment, error) {
	var payments []payment.Payment
	if err := r.db.Preload("Ticket").Find(&payments).Error; err != nil {
		return nil, err
	}
	return payments, nil
}
