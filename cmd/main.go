package main

import (
	"airline-management-system/config"
	customMiddleware "airline-management-system/internal/middleware"
	"airline-management-system/internal/router"
	"log"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	config.LoadEnv()
	db := config.InitDB()

	e := echo.New()

	// ✅ Register custom validator
	e.Validator = customMiddleware.NewValidator()

	// ✅ ใช้แค่ Logger ตัวเดียวของเรา
	e.Use(customMiddleware.ColoredLoggerMiddleware)

	// ✅ Panic-safe
	e.Use(middleware.Recover())

	// Setup routes
	router.SetupRoutes(e, db)

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: strings.Split(config.GetEnvDefault("CORS_ALLOWED_ORIGINS", "http://localhost:5173"), ","),
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders: []string{"Authorization", "Content-Type"},
	}))

	port := config.GetEnv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("Server started on port " + port)
	e.Logger.Fatal(e.Start(":" + port))
}
