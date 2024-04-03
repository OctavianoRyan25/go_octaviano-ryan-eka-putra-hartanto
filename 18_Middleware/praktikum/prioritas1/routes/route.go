package routes

import (
	"prioritas1gorm/controllers"
	"prioritas1gorm/middleware"

	gormMiddleware "github.com/labstack/echo/v4/middleware"

	"github.com/labstack/echo/v4"
)

func Init(e *echo.Echo) {
	//Routes

	e.POST("/register", controllers.Register)
	e.POST("/login", controllers.Login)

	group := e.Group("/api/v1")
	{
		group.Use(middleware.Authentication())
		group.Use(gormMiddleware.Logger())
		group.POST("/packages", controllers.AddPackage)
		group.GET("/packages", controllers.FindPackage)
		group.GET("/packages/:id", controllers.FindPackageById)
		group.PUT("/packages/:id", controllers.UpdatePackage)
		group.DELETE("/packages/:id", controllers.DeletePackage)
	}

}
