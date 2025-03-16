package router

import (
    "airline-management-system/internal/handler"
    "github.com/labstack/echo/v4"
    "gorm.io/gorm"
)

func SetupRoutes(e *echo.Echo, db *gorm.DB) {
    authHandler := handler.NewAuthHandler(nil) // ยังไม่มี UserService

    api := e.Group("/api")

    // Authentication Routes
    api.POST("/register", authHandler.Register)
}
