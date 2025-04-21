package handler

import (
	"airline-management-system/internal/service"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
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

// func (h *MaintenanceHandler) CreateMaintenanceLog(c echo.Context) error {
// 	var req maintenance.CreateMaintenanceLogRequest

// 	if err := c.Bind(&req); err != nil {
// 		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid request"})
// 	}
// 	if err := c.Validate(&req); err != nil {
// 		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
// 	}

// 	// // newLog, err := h.maintenanceService.CreateLog(&req)
// 	// if err != nil {
// 	// 	return c.JSON(http.StatusInternalServerError, echo.Map{"error": "failed to create maintenance log"})
// 	// }

// 	return c.JSON(http.StatusCreated, echo.Map{
// 		"message": "maintenance log created successfully",
// 		// "log_id":  newLog.LogID,
// 	})
// }

// func (h *MaintenanceHandler) GetMaintenanceLogDetail(c echo.Context) error {
// 	idParam := c.Param("id")
// 	idUint, err := strconv.ParseUint(idParam, 10, 64)
// 	if err != nil {
// 		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid log ID"})
// 	}

// 	log, err := h.maintenanceService.GetLogByID(uint(idUint))
// 	if err != nil {
// 		return c.JSON(http.StatusNotFound, echo.Map{"error": "maintenance log not found"})
// 	}

// 	return c.JSON(http.StatusOK, log)
// }

// func (h *MaintenanceHandler) UpdateMaintenanceLog(c echo.Context) error {
// 	idParam := c.Param("id")
// 	idUint, err := strconv.ParseUint(idParam, 10, 64)
// 	if err != nil {
// 		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid log ID"})
// 	}

// 	var req maintenance.UpdateMaintenanceLogRequest
// 	if err := c.Bind(&req); err != nil {
// 		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid request"})
// 	}

// 	err = h.maintenanceService.UpdateLog(uint(idUint), &req)
// 	if err != nil {
// 		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
// 	}

// 	return c.JSON(http.StatusOK, echo.Map{
// 		"message": "maintenance log updated successfully",
// 	})
// }

// func (h *MaintenanceHandler) DeleteMaintenanceLog(c echo.Context) error {
// 	idParam := c.Param("id")
// 	idUint, err := strconv.ParseUint(idParam, 10, 64)
// 	if err != nil {
// 		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid log ID"})
// 	}

// 	err = h.maintenanceService.DeleteLogByID(uint(idUint))
// 	if err != nil {
// 		if err.Error() == "maintenance log not found" {
// 			return c.JSON(http.StatusNotFound, echo.Map{"error": err.Error()})
// 		}
// 		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "failed to delete maintenance log"})
// 	}

// 	return c.JSON(http.StatusOK, echo.Map{
// 		"message": "maintenance log deleted successfully",
// 	})
// }
