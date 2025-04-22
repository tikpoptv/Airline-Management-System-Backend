package handler

import (
	"net/http"

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
