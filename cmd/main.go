package main

import (
	"airline-management-system/config"
	"airline-management-system/internal/container"
	customMiddleware "airline-management-system/internal/middleware"
	"airline-management-system/internal/router"
	"log"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Load ENV
	config.LoadEnv()

	// Connect Database
	db := config.InitDB()

	// Create Echo instance
	e := echo.New()

	// Register custom validator
	e.Validator = customMiddleware.NewValidator()

	// ใช้ Colored Logger
	e.Use(customMiddleware.ColoredLoggerMiddleware)

	// Panic-safe
	e.Use(middleware.Recover())

	// Setup CORS
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: strings.Split(config.GetEnvDefault("CORS_ALLOWED_ORIGINS", "http://localhost:5173"), ","),
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders: []string{"Authorization", "Content-Type"},
	}))

	// Create Container
	c := container.NewContainer(db)

	// Setup Routes
	router.SetupRoutes(e, c)

	// Start server
	port := config.GetEnv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("Server started on port " + port)
	e.Logger.Fatal(e.Start(":" + port))
}
