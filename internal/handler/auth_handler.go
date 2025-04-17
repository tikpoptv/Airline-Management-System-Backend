package handler

import (
	"net/http"

	"airline-management-system/internal/models"
	"airline-management-system/internal/service"

	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

type AuthHandler struct {
	userService *service.UserService
}

func NewAuthHandler(userService *service.UserService) *AuthHandler {
	return &AuthHandler{userService}
}

// สำหรับผู้โดยสาร register เอง
func (h *AuthHandler) RegisterPassenger(c echo.Context) error {
	var req models.RegisterRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid request"})
	}
	user, err := h.userService.RegisterPassenger(&req)
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

func (h *AuthHandler) Login(c echo.Context) error {
	var req models.LoginRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid request"})
	}

	user, err := h.userService.AuthenticateUser(req.Username, req.Password)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{"error": "invalid username or password"})
	}

	// Generate JWT Token
	claims := jwt.MapClaims{
		"user_id": user.ID,
		"role":    user.Role,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "failed to generate token"})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"token": tokenStr,
		"user": echo.Map{
			"user_id":  user.ID,
			"username": user.Username,
			"role":     user.Role,
		},
	})
}
