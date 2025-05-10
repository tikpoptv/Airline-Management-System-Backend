package handler

import (
	"airline-management-system/internal/service"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"

	routeModel "airline-management-system/internal/models/route"
)

type RouteHandler struct {
	routeService *service.RouteService
}

func NewRouteHandler(service *service.RouteService) *RouteHandler {
	return &RouteHandler{routeService: service}
}

func (h *RouteHandler) ListRoutes(c echo.Context) error {
	routes, err := h.routeService.GetAllRoutes()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "failed to fetch routes"})
	}
	return c.JSON(http.StatusOK, routes)
}

func (h *RouteHandler) CreateRoute(c echo.Context) error {
	var req routeModel.CreateRouteRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid request"})
	}
	if err := c.Validate(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	newRoute, err := h.routeService.CreateRoute(&req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, newRoute)
}

func (h *RouteHandler) UpdateRouteStatus(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid route id"})
	}

	var req routeModel.UpdateRouteStatusRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid request"})
	}
	if err := c.Validate(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	if err := h.routeService.UpdateRouteStatus(uint(id), &req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, echo.Map{"message": "route status updated successfully"})
}
