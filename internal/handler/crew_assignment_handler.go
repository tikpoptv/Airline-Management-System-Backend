package handler

import (
	"net/http"

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
