package router

import (
	"airline-management-system/internal/handler"
	"airline-management-system/internal/middleware"
	"airline-management-system/internal/repository"
	"airline-management-system/internal/service"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func SetupRoutes(e *echo.Echo, db *gorm.DB) {
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	authHandler := handler.NewAuthHandler(userService)
	adminUserHandler := handler.NewAdminUserHandler(userService)
	userHandler := handler.NewUserHandler(userService)

	// Main API group
	api := e.Group("/api")
	api.POST("/auth/register", authHandler.RegisterPassenger)
	api.POST("/auth/login", authHandler.Login)

	api.GET("/users/me", userHandler.GetMyProfile, middleware.JWTMiddleware)

	admin := api.Group("/users")
	admin.Use(middleware.JWTMiddleware)
	admin.Use(middleware.RequireRole("admin"))

	admin.POST("", adminUserHandler.CreateUserByAdmin)
	admin.GET("/:id", userHandler.GetUserProfile)
}
