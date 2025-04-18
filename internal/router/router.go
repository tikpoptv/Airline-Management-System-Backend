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

	// Aircraft Handler
	aircraftRepo := repository.NewAircraftRepository(db)
	aircraftService := service.NewAircraftService(aircraftRepo)
	aircraftHandler := handler.NewAircraftHandler(aircraftService)

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

	// Aircraft Routes (admin only)
	aircraftGroup := api.Group("/aircrafts")
	aircraftGroup.Use(middleware.JWTMiddleware)
	aircraftGroup.Use(middleware.RequireRole("admin"))

	aircraftGroup.GET("", aircraftHandler.ListAircraft)
	aircraftGroup.POST("", aircraftHandler.CreateAircraft)
	aircraftGroup.GET("/:id", aircraftHandler.GetAircraftDetail)
	aircraftGroup.PUT("/:id", aircraftHandler.UpdateAircraft)
	aircraftGroup.DELETE("/:id", aircraftHandler.DeleteAircraft)

}
