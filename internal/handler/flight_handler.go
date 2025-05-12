package handler

import (
	"airline-management-system/internal/service"
	"errors"
	"net/http"
	"strconv"
	"strings"

	"gorm.io/gorm"

	assignModel "airline-management-system/internal/models/assignment"
	flightModel "airline-management-system/internal/models/flight"

	"github.com/labstack/echo/v4"
)

type FlightHandler struct {
	flightService     *service.FlightService
	assignmentService *service.FlightAssignmentService
}

func NewFlightHandler(flightService *service.FlightService, assignmentService *service.FlightAssignmentService) *FlightHandler {
	return &FlightHandler{
		flightService:     flightService,
		assignmentService: assignmentService,
	}
}

func (h *FlightHandler) ListFlights(c echo.Context) error {
	flights, err := h.flightService.GetAllFlights()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "failed to fetch flights"})
	}
	return c.JSON(http.StatusOK, flights)
}

func (h *FlightHandler) CreateFlight(c echo.Context) error {
	var req flightModel.CreateFlightRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid request"})
	}

	if err := c.Validate(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	newFlight, err := h.flightService.CreateFlight(&req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, newFlight)
}

func (h *FlightHandler) GetFlightDetail(c echo.Context) error {
	idStr := c.Param("id")
	flightID, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid flight id"})
	}

	result, err := h.flightService.GetFlightByID(uint(flightID))
	if err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "flight not found"})
	}

	return c.JSON(http.StatusOK, result)
}

func (h *FlightHandler) UpdateFlight(c echo.Context) error {
	idStr := c.Param("id")
	flightID, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid flight id"})
	}

	var req flightModel.UpdateFlightRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid request body"})
	}

	updatedFlight, err := h.flightService.UpdateFlight(uint(flightID), &req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, updatedFlight)
}

func (h *FlightHandler) UpdateFlightDetails(c echo.Context) error {
	idStr := c.Param("id")
	flightID, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid flight id"})
	}

	var req flightModel.UpdateFlightDetailsRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid request body"})
	}

	updatedFlight, err := h.flightService.UpdateFlightDetails(uint(flightID), &req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, updatedFlight)
}

func (h *FlightHandler) DeleteFlight(c echo.Context) error {
	idStr := c.Param("id")
	flightID, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid flight id"})
	}

	err = h.flightService.DeleteFlight(uint(flightID))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusNotFound, echo.Map{"error": "flight not found"})
		}
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, echo.Map{"message": "flight deleted successfully"})
}

func (h *FlightHandler) AssignCrewToFlight(c echo.Context) error {
	flightIDParam := c.Param("flight_id")
	flightID, err := strconv.ParseUint(flightIDParam, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid flight ID"})
	}

	var req assignModel.AssignCrewRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid request"})
	}
	if err := c.Validate(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	result, err := h.assignmentService.AssignCrewToFlight(uint(flightID), &req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, echo.Map{
		"flight_id":      result.FlightID,
		"crew_id":        result.CrewID,
		"role_in_flight": result.RoleInFlight,
		"message":        "Crew assigned to flight successfully",
	})
}

func (h *FlightHandler) GetFlightCrewList(c echo.Context) error {
	flightIDParam := c.Param("flight_id")
	flightID, err := strconv.ParseUint(flightIDParam, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid flight ID"})
	}

	crewList, err := h.assignmentService.GetCrewByFlightID(uint(flightID))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "failed to fetch crew list"})
	}

	return c.JSON(http.StatusOK, crewList)
}

func (h *FlightHandler) GetFlightsByAircraftID(c echo.Context) error {
	idParam := c.Param("aircraft_id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid aircraft ID"})
	}

	flights, err := h.flightService.GetFlightsByAircraftID(uint(id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "failed to retrieve flights"})
	}

	return c.JSON(http.StatusOK, flights)
}

func (h *FlightHandler) GetFlightCrewInfo(c echo.Context) error {
	// Get flight ID from URL parameter
	flightIDStr := c.Param("flight_id")
	flightID, err := strconv.ParseUint(flightIDStr, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid flight ID"})
	}

	// Get user ID and role from JWT token
	userIDRaw := c.Get("user_id")
	userIDFloat, ok := userIDRaw.(float64)
	if !ok {
		return c.JSON(http.StatusUnauthorized, echo.Map{"error": "unauthorized"})
	}
	userID := uint(userIDFloat)

	// Get crew information
	crewInfo, err := h.assignmentService.GetFlightCrewInfo(uint(flightID), userID)
	if err != nil {
		if err.Error() == "unauthorized access" {
			return c.JSON(http.StatusForbidden, echo.Map{"error": err.Error()})
		}
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "failed to fetch crew information"})
	}

	return c.JSON(http.StatusOK, crewInfo)
}

func (h *FlightHandler) GetTodayFlights(c echo.Context) error {
	// Get status from query parameter
	status := c.QueryParam("status")

	flights, err := h.flightService.GetTodayFlights(status)
	if err != nil {
		if strings.Contains(err.Error(), "invalid status") {
			return c.JSON(http.StatusBadRequest, echo.Map{
				"error": err.Error(),
			})
		}
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": "failed to fetch today's flights",
		})
	}

	if len(flights) == 0 {
		return c.JSON(http.StatusOK, echo.Map{
			"message": "no flights found for the specified criteria",
			"flights": []interface{}{},
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"total_flights": len(flights),
		"flights":       flights,
	})
}
