package handler

import (
	"airline-management-system/internal/service"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

type MaintenanceStatsHandler struct {
	statsService *service.MaintenanceStatsService
}

func NewMaintenanceStatsHandler(statsService *service.MaintenanceStatsService) *MaintenanceStatsHandler {
	return &MaintenanceStatsHandler{statsService}
}

func (h *MaintenanceStatsHandler) GetMaintenanceStats(c echo.Context) error {
	stats, err := h.statsService.GetMaintenanceStats()
	if err != nil {
		errMsg := "failed to fetch maintenance stats"
		if strings.Contains(err.Error(), "error counting") {
			errMsg = "failed to count maintenance records"
		} else if strings.Contains(err.Error(), "error fetching today's maintenance") {
			errMsg = "failed to fetch today's maintenance records"
		}

		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error":   errMsg,
			"details": err.Error(),
		})
	}

	if stats == nil {
		return c.JSON(http.StatusNotFound, echo.Map{
			"error": "no maintenance records found",
		})
	}

	return c.JSON(http.StatusOK, stats)
}
