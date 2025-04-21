package repository

import (
	"airline-management-system/internal/models/payment"
	"errors"

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

func (r *PaymentRepository) GetPaymentByID(id uint) (*payment.Payment, error) {
	var p payment.Payment
	if err := r.db.Preload("Ticket").First(&p, "payment_id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &p, nil
}
