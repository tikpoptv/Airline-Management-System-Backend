package handler

import (
	"net/http"
	"strconv"

	"airline-management-system/internal/service"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	userService *service.UserService
}

func NewUserHandler(userService *service.UserService) *UserHandler {
	return &UserHandler{userService}
}

func (h *UserHandler) GetMyProfile(c echo.Context) error {
	userID := uint(c.Get("user_id").(float64))

	user, err := h.userService.GetUserProfile(userID)
	if err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "user not found"})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"user_id":    user.ID,
		"username":   user.Username,
		"email":      user.Email,
		"role":       user.Role,
		"is_active":  user.IsActive,
		"created_at": user.CreatedAt,
		"updated_at": user.UpdatedAt,
	})
}

func (h *UserHandler) GetUserProfile(c echo.Context) error {
	idParam := c.Param("id")
	userID, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid user id"})
	}

	user, err := h.userService.GetUserProfile(uint(userID))
	if err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "user not found"})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"user_id":    user.ID,
		"username":   user.Username,
		"email":      user.Email,
		"role":       user.Role,
		"is_active":  user.IsActive,
		"created_at": user.CreatedAt,
		"updated_at": user.UpdatedAt,
	})
}
