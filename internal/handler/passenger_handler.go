package handler

import (
	"net/http"
	"strconv"

	"airline-management-system/internal/service"

	"github.com/labstack/echo/v4"
)

type PassengerHandler struct {
	service *service.PassengerService
}

func NewPassengerHandler(service *service.PassengerService) *PassengerHandler {
	return &PassengerHandler{service}
}

func (h *PassengerHandler) GetFlightPassengers(c echo.Context) error {
	// Get flight ID from URL parameter
	flightIDStr := c.Param("flight_id")
	flightID, err := strconv.ParseUint(flightIDStr, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": "invalid flight ID",
		})
	}

	// Get passenger list
	passengers, err := h.service.GetPassengersByFlightID(uint(flightID))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error":  "failed to fetch passenger list",
			"detail": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, passengers)
}

func (h *PassengerHandler) GetPassengerByID(c echo.Context) error {
	// Get passenger ID from URL parameter
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": "invalid passenger ID",
		})
	}

	// Get passenger details
	passenger, err := h.service.GetPassengerByID(uint(id))
	if err != nil {
		if err.Error() == "passenger not found" {
			return c.JSON(http.StatusNotFound, echo.Map{
				"error": "passenger not found",
			})
		}
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error":  "failed to fetch passenger details",
			"detail": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, passenger)
}
