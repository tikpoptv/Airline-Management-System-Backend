package handler

import (
	"airline-management-system/internal/service"
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
)

type AirlineOwnerHandler struct {
	service service.AirlineOwnerService
}

func NewAirlineOwnerHandler(service service.AirlineOwnerService) *AirlineOwnerHandler {
	return &AirlineOwnerHandler{service: service}
}

func (h *AirlineOwnerHandler) GetAirlineOwners(c echo.Context) error {
	ctx := context.Background()

	owners, err := h.service.GetAllAirlineOwners(ctx)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, owners)
}
