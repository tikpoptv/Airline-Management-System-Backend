package handler

import (
	"net/http"
	"strconv"

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

func (h *PaymentHandler) GetPaymentDetail(c echo.Context) error {
	idParam := c.Param("id")
	idUint, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid payment ID"})
	}

	paymentDetail, err := h.paymentService.GetPaymentByID(uint(idUint))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "failed to retrieve payment"})
	}
	if paymentDetail == nil {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "payment not found"})
	}

	return c.JSON(http.StatusOK, paymentDetail)
}
