package main

import (
	"prioritas1gorm/configs"
	"prioritas1gorm/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	configs.ConnectDB()
	e := echo.New()
	routes.Init(e)
	e.Logger.Fatal(e.Start(":8080"))
}
