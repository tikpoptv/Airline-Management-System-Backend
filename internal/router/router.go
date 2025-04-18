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

	// Flight
	flightRepo := repository.NewFlightRepository(db)
	flightService := service.NewFlightService(flightRepo)
	flightHandler := handler.NewFlightHandler(flightService)

	// Route
	routeRepo := repository.NewRouteRepository(db)
	routeService := service.NewRouteService(routeRepo)
	routeHandler := handler.NewRouteHandler(routeService)

	// Airport
	airportRepo := repository.NewAirportRepository(db)
	airportService := service.NewAirportService(airportRepo)
	airportHandler := handler.NewAirportHandler(airportService)

	// Crew
	crewRepo := repository.NewCrewRepository(db)
	crewService := service.NewCrewService(crewRepo)
	crewHandler := handler.NewCrewHandler(crewService)

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

	// Flight Routes (admin only)
	flightGroup := api.Group("/flights")
	flightGroup.Use(middleware.JWTMiddleware)
	flightGroup.Use(middleware.RequireRole("admin"))

	flightGroup.GET("", flightHandler.ListFlights)
	flightGroup.POST("", flightHandler.CreateFlight)
	flightGroup.GET("/:id", flightHandler.GetFlightDetail)
	flightGroup.PUT("/:id", flightHandler.UpdateFlight)
	flightGroup.PUT("/:id/details", flightHandler.UpdateFlightDetails)
	flightGroup.DELETE("/:id", flightHandler.DeleteFlight)

	// Route Routes (admin only)
	routeGroup := api.Group("/routes")
	routeGroup.Use(middleware.JWTMiddleware)
	routeGroup.Use(middleware.RequireRole("admin"))

	routeGroup.GET("", routeHandler.ListRoutes)
	routeGroup.POST("", routeHandler.CreateRoute)

	// Route Details (admin only)
	airportGroup := api.Group("/airports")
	airportGroup.Use(middleware.JWTMiddleware)
	airportGroup.Use(middleware.RequireRole("admin"))

	airportGroup.GET("", airportHandler.ListAirports)
	airportGroup.POST("", airportHandler.CreateAirport)

	// Crew Routes (admin only)
	crewGroup := api.Group("/crew")
	crewGroup.Use(middleware.JWTMiddleware)
	crewGroup.Use(middleware.RequireRole("admin"))

	crewGroup.GET("", crewHandler.ListCrew)
	crewGroup.POST("", crewHandler.CreateCrew)

}
