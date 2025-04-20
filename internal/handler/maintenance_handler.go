package handler

import (
	"airline-management-system/internal/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

type MaintenanceHandler struct {
	maintenanceService *service.MaintenanceService
}

func NewMaintenanceHandler(service *service.MaintenanceService) *MaintenanceHandler {
	return &MaintenanceHandler{maintenanceService: service}
}

func (h *MaintenanceHandler) ListMaintenanceLogs(c echo.Context) error {
	logs, err := h.maintenanceService.GetAllLogs()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": "failed to fetch maintenance logs",
		})
	}
	return c.JSON(http.StatusOK, logs)
}
