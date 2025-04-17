package middleware

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func RequireRole(requiredRole string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			userRole, ok := c.Get("user_role").(string)
			if !ok || userRole != requiredRole {
				return c.JSON(http.StatusForbidden, echo.Map{"error": "forbidden: requires " + requiredRole + " role"})
			}
			return next(c)
		}
	}
}
