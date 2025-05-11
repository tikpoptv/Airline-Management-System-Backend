package handler

import (
	"airline-management-system/internal/service"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	"airline-management-system/internal/models/crew"
	crewModel "airline-management-system/internal/models/crew"
)

type CrewHandler struct {
	crewService *service.CrewService
}

func NewCrewHandler(service *service.CrewService) *CrewHandler {
	return &CrewHandler{crewService: service}
}

func (h *CrewHandler) ListCrew(c echo.Context) error {
	crews, err := h.crewService.GetAllCrew()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "failed to fetch crew"})
	}
	return c.JSON(http.StatusOK, crews)

}

func (h *CrewHandler) CreateCrew(c echo.Context) error {
	var req crewModel.CreateCrewRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid request"})
	}
	if err := c.Validate(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	newCrew, err := h.crewService.CreateCrew(&req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, newCrew)
}

func (h *CrewHandler) GetCrewDetail(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid crew ID"})
	}

	crew, err := h.crewService.GetCrewByID(uint(id))
	if err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "crew not found"})
	}

	return c.JSON(http.StatusOK, crew)
}

func (h *CrewHandler) UpdateCrew(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid crew ID"})
	}

	var req crewModel.UpdateCrewRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid request"})
	}

	if err := h.crewService.UpdateCrew(uint(id), &req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, echo.Map{"message": "crew updated successfully"})
}

func (h *CrewHandler) DeleteCrew(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid crew ID"})
	}

	ok, err := h.crewService.DeleteCrew(uint(id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "internal error"})
	}
	if !ok {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "crew not found"})
	}

	return c.JSON(http.StatusOK, echo.Map{"message": "crew deleted successfully"})
}

func (h *CrewHandler) GetCrewFlightHours(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid crew ID"})
	}

	result, err := h.crewService.GetCrewFlightHours(uint(id))
	if err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "crew not found"})
	}

	return c.JSON(http.StatusOK, result)
}

func (h *CrewHandler) UpdateMyCrewProfile(c echo.Context) error {
	// ดึง user_id จาก JWT context
	userIDRaw := c.Get("user_id")
	userIDFloat, ok := userIDRaw.(float64)
	if !ok {
		return c.JSON(http.StatusUnauthorized, echo.Map{"error": "unauthorized"})
	}
	userID := uint(userIDFloat)

	// รับข้อมูลจาก body
	var req crew.UpdateCrewProfileRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid request data"})
	}
	if err := c.Validate(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	// เรียก service เพื่ออัปเดต
	if err := h.crewService.UpdateCrewProfileFromUser(userID, &req); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusNotFound, echo.Map{"error": "crew not found"})
		}
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "failed to update profile"})
	}

	return c.JSON(http.StatusOK, echo.Map{"message": "crew profile updated successfully"})
}

func (h *CrewHandler) GetAvailableCrewsForFlight(c echo.Context) error {
	// Get flight ID from path parameter
	flightIDStr := c.Param("flight_id")
	flightID, err := strconv.ParseUint(flightIDStr, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": "invalid flight ID",
		})
	}

	crews, err := h.crewService.GetAvailableCrewsForFlight(uint(flightID))
	if err != nil {
		switch {
		case err.Error() == "flight not found":
			return c.JSON(http.StatusNotFound, echo.Map{
				"error": "flight not found",
			})
		case strings.Contains(err.Error(), "invalid departure time format"):
			return c.JSON(http.StatusBadRequest, echo.Map{
				"error": err.Error(),
			})
		case strings.Contains(err.Error(), "flight departure time must be in the future"):
			return c.JSON(http.StatusBadRequest, echo.Map{
				"error": err.Error(),
			})
		default:
			return c.JSON(http.StatusInternalServerError, echo.Map{
				"error": fmt.Sprintf("failed to fetch available crews: %v", err),
			})
		}
	}

	return c.JSON(http.StatusOK, crews)
}

func (h *CrewHandler) GetMyCrewProfile(c echo.Context) error {
	// ดึง user_id จาก JWT context
	userIDRaw := c.Get("user_id")
	userIDFloat, ok := userIDRaw.(float64)
	if !ok {
		return c.JSON(http.StatusUnauthorized, echo.Map{"error": "unauthorized"})
	}
	userID := uint(userIDFloat)

	// เรียก service เพื่อดึงข้อมูล
	crewProfile, err := h.crewService.GetCrewProfileByUserID(userID)
	if err != nil {
		if err.Error() == "crew not found" {
			return c.JSON(http.StatusNotFound, echo.Map{
				"error":   "crew profile not found",
				"message": "No crew profile found for this user. This might be because you are not registered as a crew member.",
			})
		}
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "failed to retrieve crew profile"})
	}

	return c.JSON(http.StatusOK, crewProfile)
}
