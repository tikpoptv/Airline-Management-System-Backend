package handler

import (
	"airline-management-system/internal/service"
	"net/http"

	"github.com/labstack/echo/v4"
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
