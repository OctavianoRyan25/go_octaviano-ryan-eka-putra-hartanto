package wishlist

import (
	"clean-archi/entities"

	"gorm.io/gorm"
)

type WishlistRepo struct {
	DB *gorm.DB
}

func NewWishlistRepo(db *gorm.DB) *WishlistRepo {
	return &WishlistRepo{DB: db}
}

func (wr *WishlistRepo) Create(wishlist *entities.Wishlist) error {
	wishlistDb := fromUseCase(wishlist)
	result := wr.DB.Create(&wishlistDb)
	if result.Error != nil {
		return result.Error
	}
	wishlist = wishlistDb.toUseCase()
	return nil
}

func (wr *WishlistRepo) GetAll(wishlist *entities.Wishlist) error {
	wishlistDb := fromUseCase(wishlist)
	result := wr.DB.Find(&wishlistDb)
	if result.Error != nil {
		return result.Error
	}
	wishlist = wishlistDb.toUseCase()
	return nil
}
