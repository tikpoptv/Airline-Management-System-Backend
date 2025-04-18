package handler

import (
	airportModel "airline-management-system/internal/models/airport"
	"airline-management-system/internal/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

type AirportHandler struct {
	airportService *service.AirportService
}

func NewAirportHandler(service *service.AirportService) *AirportHandler {
	return &AirportHandler{airportService: service}
}

func (h *AirportHandler) ListAirports(c echo.Context) error {
	airports, err := h.airportService.GetAllAirports()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "failed to fetch airports"})
	}
	return c.JSON(http.StatusOK, airports)
}

func (h *AirportHandler) CreateAirport(c echo.Context) error {
	var req airportModel.CreateAirportRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid request"})
	}
	if err := c.Validate(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	newAirport, err := h.airportService.CreateAirport(&req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, newAirport)
}
