package middleware

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func RequireRole(allowedRoles ...string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			userRole, ok := c.Get("user_role").(string)
			if !ok {
				return c.JSON(http.StatusForbidden, echo.Map{"error": "forbidden: missing role"})
			}

			for _, role := range allowedRoles {
				if userRole == role {
					return next(c)
				}
			}

			return c.JSON(http.StatusForbidden, echo.Map{
				"error": "forbidden: requires one of [" + joinRoles(allowedRoles) + "]",
			})
		}
	}
}

func joinRoles(roles []string) string {
	result := ""
	for i, r := range roles {
		if i > 0 {
			result += ", "
		}
		result += r
	}
	return result
}
