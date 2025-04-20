package service

import (
	"airline-management-system/internal/models/payment"
	"airline-management-system/internal/repository"
)

type PaymentService struct {
	repo *repository.PaymentRepository
}

func NewPaymentService(repo *repository.PaymentRepository) *PaymentService {
	return &PaymentService{repo}
}

func (s *PaymentService) GetAllPayments() ([]payment.Payment, error) {
	return s.repo.GetAllPayments()
}
