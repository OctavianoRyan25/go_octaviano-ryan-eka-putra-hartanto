package request

import "clean-archi/entities"

type WishlistCreateRequest struct {
	Title      string `json:"title"`
	IsAchieved bool   `json:"is_achieved"`
}

func (w *WishlistCreateRequest) ToEntities() *entities.Wishlist {
	return &entities.Wishlist{
		Title:      w.Title,
		IsAchieved: w.IsAchieved,
	}
}
