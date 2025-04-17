package handler

import (
	"net/http"

	"airline-management-system/internal/models"
	"airline-management-system/internal/service"

	"github.com/labstack/echo/v4"
)

type AdminUserHandler struct {
	userService *service.UserService
}

func NewAdminUserHandler(userService *service.UserService) *AdminUserHandler {
	return &AdminUserHandler{userService}
}

func (h *AdminUserHandler) CreateUserByAdmin(c echo.Context) error {
	var req models.AdminCreateUserRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid request"})
	}
	user, err := h.userService.AdminCreateUser(&req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, echo.Map{
		"user_id":  user.ID,
		"username": user.Username,
		"email":    user.Email,
		"role":     user.Role,
	})
}
