package controllers

import (
	"clean-archi/controllers/wishlist/request"
	"clean-archi/entities"
	"net/http"

	"github.com/labstack/echo/v4"
)

type WishlistController struct {
	wishListUseCase entities.UseCaseInterface
}

func (wl *WishlistController) GetAll(c echo.Context) error {
	var wishlist entities.Wishlist
	wishlistres, _ := wl.wishListUseCase.GetAll(&wishlist)
	return c.JSON(http.StatusOK, wishlistres)
}

func (wl *WishlistController) Create(c echo.Context) error {
	var wishlist request.WishlistCreateRequest
	c.Bind(&wishlist)

	wishlistres, _ := wl.wishListUseCase.Create(wishlist.ToEntities())
	return c.JSON(http.StatusOK, wishlistres)
}

func NewWishlistController(wishlistUseCase entities.UseCaseInterface) *WishlistController {
	return &WishlistController{wishListUseCase: wishlistUseCase}
}
