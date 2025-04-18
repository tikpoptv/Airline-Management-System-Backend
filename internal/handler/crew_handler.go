package handler

import (
	"airline-management-system/internal/service"
	"net/http"

	"github.com/labstack/echo/v4"

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
