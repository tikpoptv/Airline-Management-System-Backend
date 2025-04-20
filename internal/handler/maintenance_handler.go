package handler

import (
	"airline-management-system/internal/service"
	"net/http"
	"strconv"

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

func (h *MaintenanceHandler) GetMaintenanceLogDetail(c echo.Context) error {
	idParam := c.Param("id")
	idUint, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid log ID"})
	}

	log, err := h.maintenanceService.GetLogByID(uint(idUint))
	if err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "maintenance log not found"})
	}

	return c.JSON(http.StatusOK, log)
}

func (h *MaintenanceHandler) UpdateMaintenanceLog(c echo.Context) error {
	idParam := c.Param("id")
	idUint, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid log ID"})
	}

	var req maintenance.UpdateMaintenanceLogRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid request"})
	}

	err = h.maintenanceService.UpdateLog(uint(idUint), &req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "maintenance log updated successfully",
	})
}
