package handler

import (
	"net/http"
	"strconv"

	"airline-management-system/internal/models/maintenance"
	"airline-management-system/internal/service"

	"github.com/labstack/echo/v4"
)

type MaintenanceTaskHandler struct {
	service *service.MaintenanceTaskService
}

func NewMaintenanceTaskHandler(service *service.MaintenanceTaskService) *MaintenanceTaskHandler {
	return &MaintenanceTaskHandler{service: service}
}

func (h *MaintenanceTaskHandler) GetMyTasks(c echo.Context) error {
	userIDRaw := c.Get("user_id")
	userIDFloat, ok := userIDRaw.(float64)
	if !ok {
		return c.JSON(http.StatusUnauthorized, echo.Map{"error": "unauthorized"})
	}
	userID := uint(userIDFloat)

	tasks, err := h.service.GetTasksByUser(userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "failed to retrieve tasks"})
	}
	return c.JSON(http.StatusOK, tasks)
}

func (h *MaintenanceTaskHandler) UpdateTaskStatus(c echo.Context) error {
	logIDParam := c.Param("id")
	logID, err := strconv.Atoi(logIDParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid task ID"})
	}

	userIDRaw := c.Get("user_id")
	userIDFloat, ok := userIDRaw.(float64)
	if !ok {
		return c.JSON(http.StatusUnauthorized, echo.Map{"error": "unauthorized"})
	}
	userID := uint(userIDFloat)

	var req maintenance.UpdateTaskStatusRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid request format"})
	}
	if err := c.Validate(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	if err := h.service.UpdateTaskStatus(uint(logID), userID, &req); err != nil {
		if err.Error() == "you are not assigned to this task" {
			return c.JSON(http.StatusForbidden, echo.Map{"error": err.Error()})
		}
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, echo.Map{"message": "task status updated successfully"})
}
