package handler

import (
	"airline-management-system/internal/models"
	"airline-management-system/internal/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

type AuthHandler struct {
	userService *service.UserService
}

func NewAuthHandler(userService *service.UserService) *AuthHandler {
	return &AuthHandler{userService}
}

func (h *AuthHandler) Register(c echo.Context) error {
	var user models.User
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid request"})
	}

	if err := h.userService.RegisterUser(&user); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to register user"})
	}

	return c.JSON(http.StatusOK, echo.Map{"message": "User registered successfully"})
}
