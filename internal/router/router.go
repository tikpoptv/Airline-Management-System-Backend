package router

import (
	"airline-management-system/internal/container"
	"airline-management-system/internal/middleware"

	"github.com/labstack/echo/v4"
)

func SetupRoutes(e *echo.Echo, c *container.Container) {
	api := e.Group("/api")

	api.POST("/auth/register", c.AuthHandler.RegisterPassenger)
	api.POST("/auth/login", c.AuthHandler.Login)
	api.GET("/users/me", c.UserHandler.GetMyProfile, middleware.JWTMiddleware)

	admin := api.Group("/users", middleware.JWTMiddleware, middleware.RequireRole("admin"))
	admin.POST("", c.AdminUserHandler.CreateUserByAdmin)
	admin.GET("/:id", c.UserHandler.GetUserProfile)

	// Aircraft Routes (admin only)
	aircraftGroup := api.Group("/aircrafts", middleware.JWTMiddleware, middleware.RequireRole("admin"))
	aircraftGroup.GET("", c.AircraftHandler.ListAircraft)
	aircraftGroup.POST("", c.AircraftHandler.CreateAircraft)
	aircraftGroup.GET("/:id", c.AircraftHandler.GetAircraftDetail)
	aircraftGroup.PUT("/:id", c.AircraftHandler.UpdateAircraft)
	aircraftGroup.DELETE("/:id", c.AircraftHandler.DeleteAircraft)
	aircraftGroup.GET("/:aircraft_id/flights", c.FlightHandler.GetFlightsByAircraftID)

	// Flight Routes (admin only)
	flightGroup := api.Group("/flights", middleware.JWTMiddleware, middleware.RequireRole("admin"))
	flightGroup.GET("", c.FlightHandler.ListFlights)
	flightGroup.POST("", c.FlightHandler.CreateFlight)
	flightGroup.GET("/:id", c.FlightHandler.GetFlightDetail)
	flightGroup.PUT("/:id", c.FlightHandler.UpdateFlight)
	flightGroup.PUT("/:id/details", c.FlightHandler.UpdateFlightDetails)
	flightGroup.DELETE("/:id", c.FlightHandler.DeleteFlight)
	flightGroup.POST("/:flight_id/assign-crew", c.FlightHandler.AssignCrewToFlight)
	flightGroup.GET("/:flight_id/crew", c.FlightHandler.GetFlightCrewList)

	// Route Routes (admin only)
	routeGroup := api.Group("/routes", middleware.JWTMiddleware, middleware.RequireRole("admin"))
	routeGroup.GET("", c.RouteHandler.ListRoutes)
	routeGroup.POST("", c.RouteHandler.CreateRoute)
	routeGroup.PUT("/:id/status", c.RouteHandler.UpdateRouteStatus)

	// Airport Routes (admin only)
	airportGroup := api.Group("/airports", middleware.JWTMiddleware, middleware.RequireRole("admin"))
	airportGroup.GET("", c.AirportHandler.ListAirports)
	airportGroup.POST("", c.AirportHandler.CreateAirport)
	airportGroup.PUT("/:id", c.AirportHandler.UpdateAirport)

	// Crew Routes (admin only)
	crewGroup := api.Group("/crew", middleware.JWTMiddleware, middleware.RequireRole("admin"))
	crewGroup.GET("", c.CrewHandler.ListCrew)
	crewGroup.POST("", c.CrewHandler.CreateCrew)
	crewGroup.GET("/:id", c.CrewHandler.GetCrewDetail)
	crewGroup.PUT("/:id", c.CrewHandler.UpdateCrew)
	crewGroup.DELETE("/:id", c.CrewHandler.DeleteCrew)
	crewGroup.GET("/:id/flight-hours", c.CrewHandler.GetCrewFlightHours)
	crewGroup.GET("/:id/schedule", c.CrewAssignmentHandler.GetCrewSchedule)

	api.GET("/crew/me/assignments",
		c.CrewAssignmentHandler.GetMyAssignedFlights,
		middleware.JWTMiddleware,
		middleware.RequireRole("crew", "maintenance"),
	)

	api.PUT("/crew/me/update-profile",
		c.CrewHandler.UpdateMyCrewProfile,
		middleware.JWTMiddleware,
		middleware.RequireRole("crew"),
	)

	// Maintenance Routes (admin, maintenance only)
	maintenanceGroup := api.Group("/maintenance-logs", middleware.JWTMiddleware, middleware.RequireRole("admin", "maintenance"))
	maintenanceGroup.GET("", c.MaintenanceHandler.ListMaintenanceLogs)
	maintenanceGroup.POST("", c.MaintenanceHandler.CreateMaintenanceLog)
	maintenanceGroup.GET("/:id", c.MaintenanceHandler.GetMaintenanceLogDetail)
	maintenanceGroup.PUT("/:id", c.MaintenanceHandler.UpdateMaintenanceLog)

	// Maintenance Task
	taskGroup := api.Group("/maintenance-tasks", middleware.JWTMiddleware, middleware.RequireRole("maintenance", "admin"))
	taskGroup.GET("/me", c.MaintenanceTaskHandler.GetMyTasks)
	taskGroup.PUT("/:id/status", c.MaintenanceTaskHandler.UpdateTaskStatus)

	// Payment Routes (admin, finance, maintenance)
	paymentGroup := api.Group("/payments", middleware.JWTMiddleware, middleware.RequireRole("admin", "finance", "maintenance"))
	paymentGroup.GET("", c.PaymentHandler.ListPayments)
	paymentGroup.GET("/:id", c.PaymentHandler.GetPaymentDetail)

	// Aircraft Model Routes (admin only)
	modelGroup := api.Group("/models", middleware.JWTMiddleware, middleware.RequireRole("admin"))

	modelGroup.GET("/aircraft", c.AircraftModelHandler.GetAircraftModels)
	modelGroup.GET("/airline", c.AirlineOwnerHandler.GetAirlineOwners)
}
