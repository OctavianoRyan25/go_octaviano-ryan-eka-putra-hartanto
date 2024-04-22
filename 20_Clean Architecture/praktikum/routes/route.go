package routes

import (
	controllers "clean-archi/controllers/wishlist"

	"github.com/labstack/echo/v4"
)

type RouteController struct {
	WishlistController *controllers.WishlistController
}

func (rc *RouteController) Init(e *echo.Echo) {
	e.POST("/wishlist", rc.WishlistController.Create)
	e.GET("/wishlist", rc.WishlistController.GetAll)
}
