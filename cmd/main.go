package main

import (
	"airline-management-system/config"
	"airline-management-system/internal/router"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	config.LoadEnv()

	db := config.InitDB()

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	router.SetupRoutes(e, db)

	port := config.GetEnv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("Server started on port " + port)
	e.Logger.Fatal(e.Start(":" + port))
}
