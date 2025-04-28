package handler

import (
	"airline-management-system/internal/service"
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
)

type AircraftModelHandler struct {
	service service.AircraftModelService
}

func NewAircraftModelHandler(service service.AircraftModelService) *AircraftModelHandler {
	return &AircraftModelHandler{service: service}
}

func (h *AircraftModelHandler) GetAircraftModels(c echo.Context) error {
	ctx := context.Background()

	models, err := h.service.GetAllAircraftModels(ctx)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, models)
}
