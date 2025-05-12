package handler

import (
	"airline-management-system/internal/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

type DashboardHandler struct {
	dashboardService *service.DashboardService
}

func NewDashboardHandler(dashboardService *service.DashboardService) *DashboardHandler {
	return &DashboardHandler{dashboardService}
}

func (h *DashboardHandler) GetDashboardStats(c echo.Context) error {
	stats, err := h.dashboardService.GetDashboardStats()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": "failed to fetch dashboard stats",
		})
	}

	return c.JSON(http.StatusOK, stats)
}
