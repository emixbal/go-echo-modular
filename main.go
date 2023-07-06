package main

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"go-echo-modular/config"
	"go-echo-modular/database"
)

func main() {
	app := echo.New()

	config.InitDB()
	db := config.DB
	database.InitMigration(db)

	app.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	app.Logger.Fatal(app.Start(":1323"))
}
