package handler

import (
	"airline-management-system/internal/models/maintenance"
	"airline-management-system/internal/service"
	"errors"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type MaintenanceHandler struct {
	maintenanceService *service.MaintenanceService
}

func NewMaintenanceHandler(service *service.MaintenanceService) *MaintenanceHandler {
	return &MaintenanceHandler{maintenanceService: service}
}

func (h *MaintenanceHandler) ListMaintenanceLogs(c echo.Context) error {
	filters := make(map[string]interface{})

	if status := c.QueryParam("status"); status != "" {
		filters["status"] = status
	}
	if assigned := c.QueryParam("assigned_to"); assigned != "" {
		if uid, err := strconv.Atoi(assigned); err == nil {
			filters["assigned_to"] = uint(uid)
		}
	}
	if aircraft := c.QueryParam("aircraft_id"); aircraft != "" {
		if aid, err := strconv.Atoi(aircraft); err == nil {
			filters["aircraft_id"] = uint(aid)
		}
	}

	logs, err := h.maintenanceService.GetAllLogs(filters)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "failed to retrieve logs"})
	}
	return c.JSON(http.StatusOK, logs)
}

func (h *MaintenanceHandler) CreateMaintenanceLog(c echo.Context) error {
	var req maintenance.CreateMaintenanceLogRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid request format"})
	}
	if err := c.Validate(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	newLog, err := h.maintenanceService.CreateLog(&req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "failed to create maintenance log"})
	}

	return c.JSON(http.StatusCreated, newLog)
}

func (h *MaintenanceHandler) GetMaintenanceLogDetail(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid maintenance log ID"})
	}

	log, err := h.maintenanceService.GetLogByID(uint(id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "failed to retrieve log"})
	}
	if log == nil {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "maintenance log not found"})
	}

	return c.JSON(http.StatusOK, log)
}

func (h *MaintenanceHandler) UpdateMaintenanceLog(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid maintenance log ID"})
	}

	var req maintenance.UpdateMaintenanceLogRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid request format"})
	}

	if err := h.maintenanceService.UpdateLogByID(uint(id), &req); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusNotFound, echo.Map{"error": "maintenance log not found"})
		}
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, echo.Map{"message": "maintenance log updated successfully"})
}
