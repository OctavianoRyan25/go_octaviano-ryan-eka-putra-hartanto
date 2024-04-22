package main

import (
	"clean-archi/configs"
	controllers "clean-archi/controllers/wishlist"
	"clean-archi/drivers/mysql"
	"clean-archi/drivers/mysql/wishlist"
	"clean-archi/routes"
	"clean-archi/usecases"

	"github.com/labstack/echo/v4"
)

func main() {
	configs.LoadEnv()
	configs.InitDB()
	db := mysql.ConnectDB(configs.InitDB())
	e := echo.New()

	wishlistRepo := wishlist.NewWishlistRepo(db)
	wishlistUseCase := usecases.NewWishlistUseCase(wishlistRepo)
	wishListController := controllers.NewWishlistController(wishlistUseCase)

	routeController := routes.RouteController{
		WishlistController: wishListController,
	}

	routeController.Init(e)
	e.Logger.Fatal(e.Start(":8080"))
}
