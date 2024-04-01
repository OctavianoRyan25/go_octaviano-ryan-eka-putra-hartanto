package routes

import (
	"prioritas1gorm/controllers"

	"github.com/labstack/echo/v4"
)

func Init(e *echo.Echo) {
	//Routes
	e.POST("/api/v1/fruits", controllers.AddFruit)
	e.GET("/api/v1/fruits", controllers.FindFruit)
	e.GET("/api/v1/fruits/:id", controllers.FindFruitById)
	e.DELETE("/api/v1/fruits/:id", controllers.DeleteFruit)
}
