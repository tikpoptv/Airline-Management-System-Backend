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

func (h *CrewAssignmentHandler) GetAssignedFlights(c echo.Context) error {
	idParam := c.Param("id")
	idUint, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid crew ID"})
	}

	assignments, err := h.service.GetAssignedFlightsByCrewID(uint(idUint))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "failed to retrieve assignments"})
	}
	if len(assignments) == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "no assignments found"})
	}

	return c.JSON(http.StatusOK, assignments)
}
