package handler

import (
	"airline-management-system/internal/service"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type DashboardHandler struct {
	service *service.DashboardService
}

func NewDashboardHandler(s *service.DashboardService) *DashboardHandler {
	return &DashboardHandler{service: s}
}

func (h *DashboardHandler) GetDashboardStats(c echo.Context) error {
	stats, err := h.service.GetDashboardStats()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": "failed to fetch dashboard stats",
		})
	}

	return c.JSON(http.StatusOK, stats)
}

func (h *DashboardHandler) GetTodayCrewSchedule(c echo.Context) error {
	limitStr := c.QueryParam("limit")
	limit := 5 // Default limit

	if limitStr != "" {
		parsedLimit, err := strconv.Atoi(limitStr)
		if err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{
				"error": "invalid limit parameter",
			})
		}
		limit = parsedLimit
	}

	schedules, err := h.service.GetTodayCrewSchedule(limit)
	if err != nil {
		fmt.Printf("Error fetching crew schedule: %v\n", err)
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": fmt.Sprintf("failed to fetch crew schedule: %v", err),
		})
	}

	if len(schedules) == 0 {
		return c.JSON(http.StatusOK, echo.Map{
			"message":   "no crew schedules found for today",
			"schedules": []interface{}{},
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"total_schedules": len(schedules),
		"schedules":       schedules,
	})
}
