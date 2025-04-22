package handler

import (
	"airline-management-system/internal/service"
	"errors"
	"net/http"
	"strconv"

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
