package handler

import (
	"net/http"

	"airline-management-system/internal/service"

	aircraftModel "airline-management-system/internal/models/aircraft"

	"github.com/labstack/echo/v4"

	"strconv"
)

type AircraftHandler struct {
	aircraftService *service.AircraftService
}

func NewAircraftHandler(service *service.AircraftService) *AircraftHandler {
	return &AircraftHandler{aircraftService: service}
}

func (h *AircraftHandler) ListAircraft(c echo.Context) error {
	aircrafts, err := h.aircraftService.GetAllAircraft()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "failed to fetch aircrafts"})
	}
	return c.JSON(http.StatusOK, aircrafts)
}

func (h *AircraftHandler) CreateAircraft(c echo.Context) error {
	var req aircraftModel.CreateAircraftRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid request"})
	}

	if err := c.Validate(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	newAircraft, err := h.aircraftService.CreateAircraft(&req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, newAircraft)
}

func (h *AircraftHandler) GetAircraftDetail(c echo.Context) error {
	idStr := c.Param("id")
	aircraftID, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid aircraft id"})
	}

	result, err := h.aircraftService.GetAircraftByID(uint(aircraftID))
	if err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "aircraft not found"})
	}

	return c.JSON(http.StatusOK, result)
}

func (h *AircraftHandler) UpdateAircraft(c echo.Context) error {
	idStr := c.Param("id")
	aircraftID, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid aircraft id"})
	}

	var req aircraftModel.UpdateAircraftRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid request body"})
	}

	updatedAircraft, err := h.aircraftService.UpdateAircraft(uint(aircraftID), &req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, updatedAircraft)
}

func (h *AircraftHandler) DeleteAircraft(c echo.Context) error {
	idStr := c.Param("id")
	aircraftID, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid aircraft id"})
	}

	if err := h.aircraftService.DeleteAircraft(uint(aircraftID)); err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "aircraft not found"})
	}

	return c.JSON(http.StatusOK, echo.Map{"message": "aircraft deleted successfully"})
}
