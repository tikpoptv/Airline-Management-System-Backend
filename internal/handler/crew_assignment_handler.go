package handler

import (
	"net/http"
	"strconv"

	"airline-management-system/internal/service"

	"github.com/labstack/echo/v4"
)

type CrewAssignmentHandler struct {
	service *service.CrewAssignmentService
}

func NewCrewAssignmentHandler(s *service.CrewAssignmentService) *CrewAssignmentHandler {
	return &CrewAssignmentHandler{service: s}
}

func (h *CrewAssignmentHandler) GetMyAssignedFlights(c echo.Context) error {
	userIDRaw := c.Get("user_id")
	userIDFloat, ok := userIDRaw.(float64)
	if !ok {
		return c.JSON(http.StatusUnauthorized, echo.Map{"error": "unauthorized"})
	}
	userID := uint(userIDFloat)

	assignments, err := h.service.GetAssignedFlightsByUserID(userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "failed to retrieve assignments"})
	}
	if len(assignments) == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "no assignments found"})
	}

	return c.JSON(http.StatusOK, assignments)
}

func (h *CrewAssignmentHandler) GetCrewSchedule(c echo.Context) error {
	idParam := c.Param("id")
	crewID, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid crew ID"})
	}

	assignments, err := h.service.GetAssignedFlightsByCrewID(uint(crewID))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "failed to retrieve crew schedule"})
	}
	if len(assignments) == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "no flight assignments found for this crew"})
	}

	return c.JSON(http.StatusOK, assignments)
}
