package wishlist

import (
	"clean-archi/entities"
	"time"

	"gorm.io/gorm"
)

type Wishlist struct {
	ID         int `gorm:"primaryKey"`
	Title      string
	IsAchieved bool
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt
}

func fromUseCase(wishlist *entities.Wishlist) *Wishlist {
	return &Wishlist{
		ID:         wishlist.ID,
		Title:      wishlist.Title,
		IsAchieved: wishlist.IsAchieved,
		CreatedAt:  wishlist.CreatedAt,
		UpdatedAt:  wishlist.UpdatedAt,
	}
}

func (wishlist *Wishlist) toUseCase() *entities.Wishlist {
	return &entities.Wishlist{
		ID:         wishlist.ID,
		Title:      wishlist.Title,
		IsAchieved: wishlist.IsAchieved,
		CreatedAt:  wishlist.CreatedAt,
		UpdatedAt:  wishlist.UpdatedAt,
	}
}
