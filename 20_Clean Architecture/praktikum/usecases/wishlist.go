package usecases

import (
	"clean-archi/entities"
	"errors"
)

type WishlistUseCase struct {
	repository entities.RepositoryInterface
}

func NewWishlistUseCase(repository entities.RepositoryInterface) *WishlistUseCase {
	return &WishlistUseCase{repository: repository}
}

func (w *WishlistUseCase) GetAll(wishlist *entities.Wishlist) (entities.Wishlist, error) {
	if wishlist.Title == "" {
		return entities.Wishlist{}, errors.New("title or IsAchieved is empty")
	}
	err := w.repository.GetAll(wishlist)

	if err != nil {
		return entities.Wishlist{}, err
	}

	return *wishlist, nil
}

func (w *WishlistUseCase) Create(wishlist *entities.Wishlist) (entities.Wishlist, error) {
	if wishlist.Title == "" {
		return entities.Wishlist{}, errors.New("title or IsAchieved is empty")
	}
	err := w.repository.Create(wishlist)

	if err != nil {
		return entities.Wishlist{}, err
	}

	return *wishlist, nil
}
