package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"go-echo-modular/config"
	"go-echo-modular/database"
	"go-echo-modular/router"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	e := echo.New()

	config.InitDB()
	db := config.DB
	database.InitMigration(db)

	router.Init(e)

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "time=${time_rfc3339}, method=${method}, uri=${uri}, status=${status}, latency=${latency_human}\n",
		Output: e.Logger.Output(),
	}))
	e.Logger.Fatal(e.Start(":1323"))
}
