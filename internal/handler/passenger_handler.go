package handler

import (
	"net/http"
	"strconv"

	"airline-management-system/internal/service"

	"github.com/labstack/echo/v4"
)

type PassengerHandler struct {
	service *service.PassengerService
}

func NewPassengerHandler(service *service.PassengerService) *PassengerHandler {
	return &PassengerHandler{service}
}

func (h *PassengerHandler) ListAllPassengers(c echo.Context) error {
	// Get pagination parameters from query
	pageStr := c.QueryParam("page")
	pageSizeStr := c.QueryParam("page_size")

	// Default values
	page := 1
	pageSize := 10

	// Parse page if provided
	if pageStr != "" {
		parsedPage, err := strconv.Atoi(pageStr)
		if err == nil && parsedPage > 0 {
			page = parsedPage
		}
	}

	// Parse page size if provided
	if pageSizeStr != "" {
		parsedPageSize, err := strconv.Atoi(pageSizeStr)
		if err == nil && parsedPageSize > 0 && parsedPageSize <= 100 {
			pageSize = parsedPageSize
		}
	}

	// Get passengers from service
	passengers, err := h.service.GetAllPassengers(page, pageSize)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error":  "failed to fetch passengers",
			"detail": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, passengers)
}

func (h *PassengerHandler) GetFlightPassengers(c echo.Context) error {
	// Get flight ID from URL parameter
	flightIDStr := c.Param("flight_id")
	flightID, err := strconv.ParseUint(flightIDStr, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": "invalid flight ID",
		})
	}

	// Get passenger list
	passengers, err := h.service.GetPassengersByFlightID(uint(flightID))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error":  "failed to fetch passenger list",
			"detail": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, passengers)
}

func (h *PassengerHandler) GetPassengerByID(c echo.Context) error {
	// Get passenger ID from URL parameter
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": "invalid passenger ID",
		})
	}

	// Get passenger details
	passenger, err := h.service.GetPassengerByID(uint(id))
	if err != nil {
		if err.Error() == "passenger not found" {
			return c.JSON(http.StatusNotFound, echo.Map{
				"error": "passenger not found",
			})
		}
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error":  "failed to fetch passenger details",
			"detail": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, passenger)
}

func (h *PassengerHandler) SearchPassengers(c echo.Context) error {
	// Get search query
	query := c.QueryParam("q")
	if query == "" {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": "search query parameter 'q' is required",
		})
	}

	// Get pagination parameters from query
	pageStr := c.QueryParam("page")
	pageSizeStr := c.QueryParam("page_size")

	// Default values
	page := 1
	pageSize := 10

	// Parse page if provided
	if pageStr != "" {
		parsedPage, err := strconv.Atoi(pageStr)
		if err == nil && parsedPage > 0 {
			page = parsedPage
		}
	}

	// Parse page size if provided
	if pageSizeStr != "" {
		parsedPageSize, err := strconv.Atoi(pageSizeStr)
		if err == nil && parsedPageSize > 0 && parsedPageSize <= 100 {
			pageSize = parsedPageSize
		}
	}

	// Search passengers
	results, err := h.service.SearchPassengers(query, page, pageSize)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error":  "failed to search passengers",
			"detail": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, results)
}
