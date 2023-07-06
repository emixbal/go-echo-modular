package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"

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

	e.Logger.Fatal(e.Start(":1323"))
}
