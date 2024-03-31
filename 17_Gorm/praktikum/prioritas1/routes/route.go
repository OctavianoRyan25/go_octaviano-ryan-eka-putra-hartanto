package routes

import (
	"prioritas1gorm/controllers"

	"github.com/labstack/echo/v4"
)

func Init(e *echo.Echo) {
	//Routes
	e.POST("/api/v1/packages", controllers.AddPackage)
	e.GET("/api/v1/packages", controllers.FindPackage)
	e.GET("/api/v1/packages/:id", controllers.FindPackageById)
	e.PUT("/api/v1/packages/:id", controllers.UpdatePackage)
	e.DELETE("/api/v1/packages/:id", controllers.DeletePackage)
}
