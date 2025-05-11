package container

import (
	"airline-management-system/internal/database"
	"airline-management-system/internal/handler"
	"airline-management-system/internal/repository"
	"airline-management-system/internal/service"

	"gorm.io/gorm"
)

type Container struct {
	// User
	UserHandler      *handler.UserHandler
	AdminUserHandler *handler.AdminUserHandler
	AuthHandler      *handler.AuthHandler

	// Aircraft
	AircraftHandler *handler.AircraftHandler

	// Flight
	FlightHandler *handler.FlightHandler

	// Route
	RouteHandler *handler.RouteHandler

	// Airport
	AirportHandler *handler.AirportHandler

	// Crew
	CrewHandler           *handler.CrewHandler
	CrewAssignmentHandler *handler.CrewAssignmentHandler

	// Maintenance
	MaintenanceHandler     *handler.MaintenanceHandler
	MaintenanceTaskHandler *handler.MaintenanceTaskHandler

	// Payment
	PaymentHandler *handler.PaymentHandler

	// AircraftModel
	AircraftModelHandler *handler.AircraftModelHandler

	// AirlineOwner
	AirlineOwnerHandler *handler.AirlineOwnerHandler

	// Passenger
	PassengerHandler *handler.PassengerHandler
}

func NewContainer(db *gorm.DB) *Container {
	dbService := database.NewDBService(db)

	// User
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	authHandler := handler.NewAuthHandler(userService)
	adminUserHandler := handler.NewAdminUserHandler(userService)
	userHandler := handler.NewUserHandler(userService)

	// Passenger
	passengerRepo := repository.NewPassengerRepository(db)
	passengerService := service.NewPassengerService(passengerRepo)
	passengerHandler := handler.NewPassengerHandler(passengerService)

	// Aircraft
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
	crewService := service.NewCrewService(crewRepo, flightRepo)
	crewHandler := handler.NewCrewHandler(crewService)

	assignmentRepo := repository.NewFlightAssignmentRepository(db)
	crewAssignmentService := service.NewCrewAssignmentService(crewRepo, assignmentRepo)
	crewAssignmentHandler := handler.NewCrewAssignmentHandler(crewAssignmentService)

	// Maintenance
	maintenanceRepo := repository.NewMaintenanceRepository(db)
	maintenanceService := service.NewMaintenanceService(maintenanceRepo)
	maintenanceHandler := handler.NewMaintenanceHandler(maintenanceService)

	// Maintenance Task
	taskRepo := repository.NewMaintenanceTaskRepository(db)
	taskService := service.NewMaintenanceTaskService(taskRepo)
	taskHandler := handler.NewMaintenanceTaskHandler(taskService)

	// Payment
	paymentRepo := repository.NewPaymentRepository(db)
	paymentService := service.NewPaymentService(paymentRepo)
	paymentHandler := handler.NewPaymentHandler(paymentService)

	// Aircraft Model
	aircraftModelRepo := repository.NewAircraftModelRepository(dbService)
	aircraftModelService := service.NewAircraftModelService(aircraftModelRepo)
	aircraftModelHandler := handler.NewAircraftModelHandler(aircraftModelService)

	// Airline Owner
	airlineOwnerRepo := repository.NewAirlineOwnerRepository(dbService)
	airlineOwnerService := service.NewAirlineOwnerService(airlineOwnerRepo)
	airlineOwnerHandler := handler.NewAirlineOwnerHandler(airlineOwnerService)

	return &Container{
		// User
		UserHandler:      userHandler,
		AdminUserHandler: adminUserHandler,
		AuthHandler:      authHandler,

		// Aircraft
		AircraftHandler: aircraftHandler,

		// Flight
		FlightHandler: flightHandler,

		// Route
		RouteHandler: routeHandler,

		// Airport
		AirportHandler: airportHandler,

		// Crew
		CrewHandler:           crewHandler,
		CrewAssignmentHandler: crewAssignmentHandler,

		// Maintenance
		MaintenanceHandler:     maintenanceHandler,
		MaintenanceTaskHandler: taskHandler,

		// Payment
		PaymentHandler: paymentHandler,

		// Aircraft Model
		AircraftModelHandler: aircraftModelHandler,
		AirlineOwnerHandler:  airlineOwnerHandler,

		// Passenger
		PassengerHandler: passengerHandler,
	}
}
