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

	flightAssignmentRepo := repository.NewFlightAssignmentRepository(db)
	flightAssignmentService := service.NewFlightAssignmentService(flightAssignmentRepo)

	flightHandler := handler.NewFlightHandler(flightService, flightAssignmentService)

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

	// Maintenance
	maintenanceRepo := repository.NewMaintenanceRepository(db)
	maintenanceService := service.NewMaintenanceService(maintenanceRepo)
	maintenanceHandler := handler.NewMaintenanceHandler(maintenanceService)

	// Payment
	paymentRepo := repository.NewPaymentRepository(db)
	paymentService := service.NewPaymentService(paymentRepo)
	paymentHandler := handler.NewPaymentHandler(paymentService)

	// Crew Assignment
	assignmentRepo := repository.NewFlightAssignmentRepository(db)
	crewAssignmentService := service.NewCrewAssignmentService(crewRepo, assignmentRepo)
	crewAssignmentHandler := handler.NewCrewAssignmentHandler(crewAssignmentService)

	// Maintenance Task
	taskRepo := repository.NewMaintenanceTaskRepository(db)
	taskService := service.NewMaintenanceTaskService(taskRepo)
	taskHandler := handler.NewMaintenanceTaskHandler(taskService)

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

	// Flight Assignment Routes (admin only)
	flightGroup.POST("/:flight_id/assign-crew", flightHandler.AssignCrewToFlight)
	flightGroup.GET("/:flight_id/crew", flightHandler.GetFlightCrewList)

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
	crewGroup.GET("/:id", crewHandler.GetCrewDetail)
	crewGroup.PUT("/:id", crewHandler.UpdateCrew)
	crewGroup.DELETE("/:id", crewHandler.DeleteCrew)
	crewGroup.GET("/:id/flight-hours", crewHandler.GetCrewFlightHours)

	api.GET("/crew/me/assignments",
		crewAssignmentHandler.GetMyAssignedFlights,
		middleware.JWTMiddleware,
		middleware.RequireRole("crew", "maintenance"),
	)

	api.PUT("/crew/me/update-profile",
		crewHandler.UpdateMyCrewProfile,
		middleware.JWTMiddleware,
		middleware.RequireRole("crew"),
	)

	// Maintenance Routes (admin only)
	maintenanceGroup := api.Group("/maintenance-logs")
	maintenanceGroup.Use(middleware.JWTMiddleware)
	maintenanceGroup.Use(middleware.RequireRole("admin", "maintenance"))

	maintenanceGroup.GET("", maintenanceHandler.ListMaintenanceLogs)
	maintenanceGroup.POST("", maintenanceHandler.CreateMaintenanceLog)
	maintenanceGroup.GET("/:id", maintenanceHandler.GetMaintenanceLogDetail)
	maintenanceGroup.PUT("/:id", maintenanceHandler.UpdateMaintenanceLog)
	// maintenanceGroup.DELETE("/:id", maintenanceHandler.DeleteMaintenanceLog)

	// Payment Routes (admin only)
	paymentGroup := api.Group("/payments")
	paymentGroup.Use(middleware.JWTMiddleware)
	paymentGroup.Use(middleware.RequireRole("admin", "finance", "maintenance"))
	paymentGroup.GET("", paymentHandler.ListPayments)
	paymentGroup.GET("/:id", paymentHandler.GetPaymentDetail)

	taskGroup := api.Group("/maintenance-tasks")
	taskGroup.Use(middleware.JWTMiddleware)
	taskGroup.Use(middleware.RequireRole("maintenance", "admin"))
	taskGroup.GET("/me", taskHandler.GetMyTasks)
	taskGroup.PUT("/:id/status", taskHandler.UpdateTaskStatus)

}
