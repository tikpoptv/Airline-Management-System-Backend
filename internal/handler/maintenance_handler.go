package handler

import (
	"airline-management-system/internal/service"
	"net/http"

	"airline-management-system/internal/models/maintenance"

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

func (h *MaintenanceHandler) CreateMaintenanceLog(c echo.Context) error {
	var req maintenance.CreateMaintenanceLogRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid request"})
	}
	if err := c.Validate(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	newLog, err := h.maintenanceService.CreateLog(&req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "failed to create maintenance log"})
	}

	return c.JSON(http.StatusCreated, echo.Map{
		"message": "maintenance log created successfully",
		"log_id":  newLog.LogID,
	})
}
