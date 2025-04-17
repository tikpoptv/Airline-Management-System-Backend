package main

import (
	"airline-management-system/config"
	customMiddleware "airline-management-system/internal/middleware" // 👈 ตั้ง alias ให้ไม่ชนกับ echo/middleware
	"airline-management-system/internal/router"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	config.LoadEnv()
	db := config.InitDB()

	e := echo.New()

	// ✅ ใช้แค่ Logger ตัวเดียวของเรา
	e.Use(customMiddleware.ColoredLoggerMiddleware)

	// ✅ Panic-safe
	e.Use(middleware.Recover())

	// Setup routes
	router.SetupRoutes(e, db)

	port := config.GetEnv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("Server started on port " + port)
	e.Logger.Fatal(e.Start(":" + port))
}
