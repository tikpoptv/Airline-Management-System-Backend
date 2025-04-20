package handler

import (
	"net/http"

	"airline-management-system/internal/service"

	"github.com/labstack/echo/v4"
)

type PaymentHandler struct {
	paymentService *service.PaymentService
}

func NewPaymentHandler(paymentService *service.PaymentService) *PaymentHandler {
	return &PaymentHandler{paymentService: paymentService}
}

func (h *PaymentHandler) ListPayments(c echo.Context) error {
	payments, err := h.paymentService.GetAllPayments()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": "failed to retrieve payments",
		})
	}
	return c.JSON(http.StatusOK, payments)
}
